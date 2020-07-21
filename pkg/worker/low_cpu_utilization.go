package worker

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/pipetail/cloudlint/internal/utils"
	"github.com/pipetail/cloudlint/pkg/awsregions"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

// WeightedAverage stores weighted average
type WeightedAverage struct {
	Value  float64 `json:"value"`
	Weight float64 `json:"weight"`
}

func getCPUUtilizationWithinRegion(client *cloudwatch.CloudWatch) *WeightedAverage {

	result := &WeightedAverage{
		Value:  0.0,
		Weight: 1,
	}

	input := cloudwatch.GetMetricDataInput{
		MetricDataQueries: []*cloudwatch.MetricDataQuery{
			// {
			// 	Id:         aws.String("e1"),
			// 	Expression: aws.String("AVG(METRICS())"),
			// 	Label:      aws.String("Expression1"),
			// },
			{
				Id: aws.String("m1"),
				MetricStat: &cloudwatch.MetricStat{
					Metric: &cloudwatch.Metric{
						Namespace:  aws.String("AWS/EC2"),
						MetricName: aws.String("CPUUtilization"),
					},
					Period: aws.Int64(60 * 60 * 24), // 1 day period
					Stat:   aws.String("Average"),
				},
			},
		},
		StartTime: utils.GetTimeTwoWeeksAgoStart(),
		EndTime:   utils.GetTimeNow(),
	}

	log.WithFields(log.Fields{
		"input": input,
	}).Info("checking lowcpuutilization")

	resp, err := client.GetMetricData(&input)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("checking lowcpuutilization")
		return nil
	}

	avg := 0.0
	sum := 0.0
	records := 0

	for _, metricdata := range resp.MetricDataResults {
		log.WithFields(log.Fields{
			"metricData": metricdata,
		}).Info("checking lowcpuutilization")

		records += len(metricdata.Values)
		for index := range metricdata.Timestamps {

			log.WithFields(log.Fields{
				"metricdata.Timestamps[index]": metricdata.Timestamps[index],
				"metricdata.Values[index]":     *metricdata.Values[index],
			}).Info("checking lowcpuutilization")
			sum += *metricdata.Values[index]
		}
	}

	avg = sum / float64(records)

	result.Value = avg
	result.Weight = float64(records)

	log.WithFields(log.Fields{
		"sum":      sum,
		"#records": records,
		"avg":      avg,
	}).Info("checking lowcpuutilization")

	return result
}

func getPricePerMonth(price ec2price) (float64, error) {
	if price.unit != "Hrs" {

		log.WithFields(log.Fields{
			"priceUnit": price.unit,
		}).Error("priceUnit is not Hrs")

		return 0, fmt.Errorf("price has wrong unit: %s", price.unit)
	}

	if price.value < 0 {
		return 0, nil
	}

	return price.value * 24 * 30, nil
}

type ec2price struct {
	value float64
	unit  string
}

func getSomeKey(m map[string]interface{}) string {

	// return the first key you find, we don't care which is that
	for k := range m {
		return k
	}
	return ""
}

func extractPrice(resp *pricing.GetProductsOutput) ec2price {

	// check if there is exactly one item in the PriceList
	if len(resp.PriceList) != 1 {
		log.WithFields(log.Fields{
			"len(resp.PriceList)": len(resp.PriceList),
		}).Error("only one item in PriceList expected")
		panic(fmt.Sprintf("%v", resp.PriceList))
	}

	onDemand := resp.PriceList[0]["terms"].(map[string]interface{})["OnDemand"]

	// we will use getSomeKey so that we don't have to use reflect
	//keys := reflect.ValueOf(onDemand).MapKeys()
	keys := onDemand.(map[string]interface{})

	// check if we can extract only one product key
	if len(keys) != 1 {
		log.WithFields(log.Fields{
			"len(keys)": len(keys),
		}).Error("only one Product Key expected")
		panic(fmt.Sprintf("%v", keys))
	}

	productCode := getSomeKey(keys)

	priceDimensions := onDemand.(map[string]interface{})[productCode].(map[string]interface{})["priceDimensions"]

	pcKeys := priceDimensions.(map[string]interface{})

	priceDimensionsKey := getSomeKey(pcKeys)

	price := priceDimensions.(map[string]interface{})[priceDimensionsKey].(map[string]interface{})["pricePerUnit"].(map[string]interface{})["USD"].(string)
	priceUnit := priceDimensions.(map[string]interface{})[priceDimensionsKey].(map[string]interface{})["unit"].(string)

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"price": price,
		}).Error("convert price to float64")
		panic(fmt.Sprintf("%v", price))
	}

	log.WithFields(log.Fields{
		"respSize": len(resp.PriceList),
		// "resp":           resp.PriceList,
		"productCode":    productCode,
		"priceDimension": priceDimensions,
		"price":          price,
		"len(keys)":      len(keys),
	}).Info("getMonthlyPriceOfInstance")

	return ec2price{priceFloat, priceUnit}

}

// this check is super naive as it checks the instances that are running RIGHT now (which might ignore any peaks or overall usage per month)
// but we check it against utilization from all of the instances
func getMonthlyPriceOfInstance(client *pricing.Pricing, machineType string, region string) float64 {

	input := pricing.GetProductsInput{
		Filters: []*pricing.Filter{
			{
				Field: aws.String("ServiceCode"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("AmazonEC2"),
			},
			{
				Field: aws.String("Location"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String(utils.GetLocationForRegion(region)),
			},
			{
				Field: aws.String("instanceType"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String(machineType),
			},
			// {
			// 	Field: aws.String("termType"),
			// 	Type:  aws.String("TERM_MATCH"),
			// 	Value: aws.String("OnDemand"),
			// },
			{
				Field: aws.String("operatingSystem"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("Linux"),
			},
			{
				Field: aws.String("preInstalledSw"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("NA"),
			},
			{
				Field: aws.String("tenancy"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("Shared"),
			},
			// {
			// 	Field: aws.String("operation"),
			// 	Type:  aws.String("TERM_MATCH"),
			// 	Value: aws.String("RunInstance"),
			// },
			{
				Field: aws.String("capacitystatus"),
				Type:  aws.String("TERM_MATCH"),
				Value: aws.String("Used"),
			},
		},
		FormatVersion: aws.String("aws_v1"),
		MaxResults:    aws.Int64(10),
	}

	// this is a workaround for a bug: https://github.com/aws/aws-sdk-go/issues/3323
	input.SetServiceCode("AmazonEC2")

	log.WithFields(log.Fields{
		"input": input,
	}).Info("getMonthlyPriceOfInstance")

	resp, err := client.GetProducts(&input)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("checking getMonthlyPriceOfInstance")
		return 0
	}

	// fmt.Printf("------------\npriceList: %#v\n\n", resp.PriceList)

	// data, err := json.MarshalIndent(resp, "", "\t")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%s\n", data)

	// var response awspricing.PricingResponse

	// json.Unmarshal([]byte(resp.GoString()), &response)

	price := extractPrice(resp)

	pricePerMonth, err := getPricePerMonth(price)

	if err != nil {
		return 0 // TODO: we should return the err!
	}

	return pricePerMonth
}

func getEC2InstancesPriceWithinRegion(ec2client *ec2.EC2, pricingClient *pricing.Pricing, region string) float64 {

	price := 0.0

	input := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending")},
			},
		},
	}

	log.WithFields(log.Fields{
		"input": input,
	}).Info("getEC2InstancesPriceWithinRegion")

	resp, err := ec2client.DescribeInstances(input)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("checking getEC2InstancesPriceWithinRegion")
		return 0
	}

	for idx, res := range resp.Reservations {

		log.WithFields(log.Fields{
			"ReservationId": *res.ReservationId,
			"#Instances":    len(res.Instances),
		}).Info("getEC2InstancesPriceWithinRegion")

		for _, inst := range resp.Reservations[idx].Instances {

			price += getMonthlyPriceOfInstance(pricingClient, *inst.InstanceType, region)
			log.WithFields(log.Fields{
				"InstanceId": *inst.InstanceId,
			}).Info("getEC2InstancesPriceWithinRegion")
		}
	}

	return price
}

func lowcpuutilization(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	auth := event.Payload.AWSAuth

	log.WithFields(log.Fields{
		"event": event,
	}).Info("checking lowcpuutilization")

	//var countDisks int64 = 0
	var totalMonthlyPrice float64 = 0
	wAvgTotal := WeightedAverage{
		Value:  0,
		Weight: 0,
	}

	regions := awsregions.GetRegions()

	pricingClient := NewPricingClient(auth)

	// see https://godoc.org/github.com/aws/aws-sdk-go/service/ec2#Region
	for _, region := range regions {

		log.WithFields(log.Fields{
			"awsRegion": region,
		}).Debug("checking lowcpuutilization in aws region")

		svc := NewCWClient(auth, region)

		wAvg := getCPUUtilizationWithinRegion(svc)

		// no cloudwatch metricdata for this region
		if wAvg.Weight == 0 {
			continue
		}

		// compute weighted average of the two averages and their weights
		if (wAvgTotal.Weight + wAvg.Weight) > 0.0 {
			wAvgTotal.Value = (wAvgTotal.Value*wAvgTotal.Weight + wAvg.Value*wAvg.Value) / (wAvgTotal.Weight + wAvg.Weight)
			wAvgTotal.Weight += wAvg.Weight
		}

		ec2client := NewEC2Client(auth, region)
		totalMonthlyPrice += getEC2InstancesPriceWithinRegion(ec2client, pricingClient, region)

		// count the price
		// TODO: ec2 client to count the price for ec2
		//totalMonthlyPrice += float64(eipCountWithinRegion) * getAddressPriceInRegion(region) * (24 * 30)
	}

	severity := checkcompleted.INFO
	if wAvgTotal.Value <= 50 { // 50% utilization
		severity = checkcompleted.WARNING
	}

	// set check details
	outputReport.Payload.Check.Severity = severity

	// impact is ( 100% minus percentageOfUtilization) * totalMonthlyPrice
	outputReport.Payload.Check.Impact = int(totalMonthlyPrice * ((100 - wAvgTotal.Value) / 100))

	log.WithFields(log.Fields{
		"checkCompleted":    outputReport,
		"totalMonthlyPrice": totalMonthlyPrice,
		"wAvgTotal.Value":   wAvgTotal.Value,
	}).Info("lowcpuutilization check finished")

	return &outputReport, nil
}
