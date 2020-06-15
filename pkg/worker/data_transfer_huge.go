package worker

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func getEgressPrice(ce *costexplorer.CostExplorer) (price float64) {

	params := &costexplorer.GetCostAndUsageInput{

		Granularity: aws.String("MONTHLY"),
		Metrics:     []*string{aws.String("UnblendedCost")},
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(getLastMonthStart()),
			End:   aws.String(getLastMonthEnd()),
		},
		Filter: &costexplorer.Expression{

			Dimensions: &costexplorer.DimensionValues{
				Key:    aws.String("USAGE_TYPE"),
				Values: []*string{aws.String("DataTransfer")},
			},
		},
		// GroupBy: []*costexplorer.GroupDefinition{
		// 	{
		// 		Type: aws.String("DIMENSION"),
		// 		Key:  aws.String("RESOURCE_ID"),
		// 	},
		// },
	}

	res, err := ce.GetCostAndUsage(params)
	if err != nil {
		log.WithFields(log.Fields{
			"costParams": params,
			"err":        err,
		}).Error("calling GetCostAndUsage returned error")
		return 0
	}

	log.WithFields(log.Fields{
		"res": res,
	}).Info("checking datatransfer_huge")

	//price = 30
	priceString := *res.ResultsByTime[len(res.ResultsByTime)-1].Total["UnblendedCost"].Amount
	price, _ = strconv.ParseFloat(priceString, 64)

	log.WithFields(log.Fields{
		"price":       price,
		"priceString": priceString,
	}).Debug("final price res.ResultsByTime")

	return price
}

func datatransferhuge(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	auth := event.Payload.AWSAuth

	log.WithFields(log.Fields{
		"event": event,
	}).Debug("checking datatransfer_huge")

	var totalMonthlyPrice float64 = 0

	// cost explorer is available only in us-east-1
	ce := NewCEClient(auth)

	price := getEgressPrice(ce)

	// count the price
	totalMonthlyPrice = price

	// TODO: make this relative to total spend
	severity := checkcompleted.INFO
	if totalMonthlyPrice > 0 {
		severity = checkcompleted.WARNING
	}

	// set check details
	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(totalMonthlyPrice)

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("datatransfer_huge check finished")

	return &outputReport, nil

}
