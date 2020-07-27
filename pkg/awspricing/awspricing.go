package awspricing

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/pricing/pricingiface"
	"github.com/pipetail/cloudlint/internal/utils"
	log "github.com/sirupsen/logrus"
)

type ec2price struct {
	value float64
	unit  string
}

func getPricePerMonth(price ec2price) (float64, error) {
	if price.value < 0 {
		return 0, nil
	}

	switch price.unit {
	case "Hrs":
		return price.value * 24 * 30, nil
	case "GB-Mo":
		return price.value, nil
	default:
		log.WithFields(log.Fields{
			"priceUnit": price.unit,
		}).Error("priceUnit is recognized")

		return 0, fmt.Errorf("price has wrong Unit: %s", price.unit)
	}
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

func getPrice(client pricingiface.PricingAPI, filters []*pricing.Filter) float64 {
	log.WithFields(log.Fields{
		"filters": filters,
	}).Info("getPriceOfValue")

	input := pricing.GetProductsInput{
		Filters:       filters,
		FormatVersion: aws.String("aws_v1"),
		MaxResults:    aws.Int64(10),
	}

	// this is a workaround for a bug: https://github.com/aws/aws-sdk-go/issues/3323
	input.SetServiceCode("AmazonEC2")

	log.WithFields(log.Fields{
		"input": input,
	}).Info("getPriceOfValue")

	resp, err := client.GetProducts(&input)

	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("checking getPriceOfValue")
		return 0
	}

	price := extractPrice(resp)

	pricePerMonth, err := getPricePerMonth(price)

	if err != nil {
		return 0 // TODO: we should return the err!
	}

	return pricePerMonth
}

// GetMonthlyPriceOfInstance check is super naive as it checks the instances that are running RIGHT now (which might ignore any peaks or overall usage per month)
// but we check it against utilization from all of the instances
func GetMonthlyPriceOfInstance(client pricingiface.PricingAPI, machineType string, region string) float64 {
	filters := []*pricing.Filter{
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
		{
			Field: aws.String("capacitystatus"),
			Type:  aws.String("TERM_MATCH"),
			Value: aws.String("Used"),
		},
	}

	return getPrice(client, filters)
}

// GetPriceOfVolume price of the volume type within the region
func GetPriceOfVolume(client pricingiface.PricingAPI, volumeType string, region string) float64 {
    filters := []*pricing.Filter{
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
            Field: aws.String("volumeType"),
            Type:  aws.String("TERM_MATCH"),
            Value: aws.String(utils.ConvertVolumeTypeForPricing(volumeType)),
        },
    }

    return getPrice(client, filters)
}
