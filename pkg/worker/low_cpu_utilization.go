package worker

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/cloudwatch"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/pricing"
    "github.com/pipetail/cloudlint/internal/utils"
    "github.com/pipetail/cloudlint/pkg/awspricing"
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

			price += awspricing.GetMonthlyPriceOfInstance(pricingClient, *inst.InstanceType, region)
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
