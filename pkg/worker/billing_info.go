package worker

import (
	"github.com/aws/aws-sdk-go/aws"

	"github.com/pipetail/cloudlint/internal/utils"

	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"

	"strconv"

	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func billingInfo(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	log.WithFields(log.Fields{
		"CheckID": event.Payload.CheckID,
	}).Info("Checked Unblended cost for all services in cost explorer")

	var amount float64

	auth := event.Payload.AWSAuth

	ce := NewCEClient(auth)

	costParams := &costexplorer.GetCostAndUsageInput{
		//Filter:      &costexplorer.Expression{},
		Granularity: aws.String("MONTHLY"),
		Metrics:     []*string{aws.String("UnblendedCost")},
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(utils.GetLastBillingPeriodStart()),
			End:   aws.String(utils.GetLastBillingPeriodEnd()),
		},
	}

	log.WithFields(log.Fields{
		"costParams": costParams,
	}).Debug("calling GetCostAndUsage")

	// Call to get detailed information on each instance
	cost, err := ce.GetCostAndUsage(costParams)
	if err != nil {
		log.WithFields(log.Fields{
			"costParams": costParams,
		}).Error("calling GetCostAndUsage returned error")
		return &outputReport, err
	}

	log.WithFields(log.Fields{
		"cost": cost,
	}).Debug("GetCostAndUsage successful:")

	sum := 0.0

	for _, element := range cost.ResultsByTime {
		//sum += strconv.Atoi(*element.Total["UnblendedCost"].Amount)
		amount := *element.Total["UnblendedCost"].Amount
		i, _ := strconv.ParseFloat(amount, 64)
		sum += i
		log.WithFields(log.Fields{
			"amount":  amount,
			"strconv": i,
			"sum":     sum,
		}).Debug("sum cost.ResultsByTime")
		//fmt.Println("Total cost for last month", amount)
	}

	amountString := *cost.ResultsByTime[len(cost.ResultsByTime)-1].Total["UnblendedCost"].Amount
	amount, _ = strconv.ParseFloat(amountString, 64)

	log.WithFields(log.Fields{
		"amount":       amount,
		"amountString": amountString,
	}).Debug("final amount cost.ResultsByTime")

	// set check details
	outputReport.Payload.Check.Severity = checkcompleted.INFO
	outputReport.Payload.Check.Impact = int(amount)

	return &outputReport, nil

}
