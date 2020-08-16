package awspricing

import (
	"encoding/json"
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

func extractPrice(resp *pricing.GetProductsOutput) ec2price {

	// check if there is exactly one item in the PriceList
	if len(resp.PriceList) != 1 {
		log.WithFields(log.Fields{
			"len(resp.PriceList)": len(resp.PriceList),
		}).Error("only one item in PriceList expected")
		panic(fmt.Sprintf("%v", resp.PriceList))
	}

	res := GetProductResponse{}
	respJ, err := json.Marshal(resp)
	if err != nil {
		panic(fmt.Sprintf("could not parse input json: %s", err))
	}

	err = json.Unmarshal(respJ, &res)
	if err != nil {
		panic(fmt.Sprintf("could not parse input json: %s", err))
	}

	onDemand := res.PriceList[0].Terms.OnDemand

	// get first key
	productCode := func(m map[string]Offering) string {
		for k := range m {
			return k
		}
		return ""
	}(onDemand)

	priceDimensions := onDemand[productCode].PriceDimensions

	// get first key
	priceDimensionsKey := func(m map[string]PriceDimension) string {
		for k := range m {
			return k
		}
		return ""
	}(priceDimensions)

	price := priceDimensions[priceDimensionsKey].PricePerUnit.USD
	priceUnit := priceDimensions[priceDimensionsKey].Unit

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.WithFields(log.Fields{
			"price": price,
		}).Error("convert price to float64")
		panic(fmt.Sprintf("%v", price))
	}

	log.WithFields(log.Fields{
		"respSize":       len(resp.PriceList),
		"productCode":    productCode,
		"priceDimension": priceDimensions,
		"price":          price,
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
			Value: aws.String("US East (N. Virginia)"),
		},
		{
			Field: aws.String("volumeType"),
			Type:  aws.String("TERM_MATCH"),
			Value: aws.String(utils.TranslateVolumeType(volumeType)),
		},
	}

	return getPrice(client, filters)
}
