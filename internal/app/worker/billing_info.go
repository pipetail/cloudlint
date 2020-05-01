package worker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

	var amount int64

	//Create new Cost Explorer client
	// authenticate to AWS
	sess := session.Must(session.NewSession())

	// creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
	// 	p.ExternalID = &externalID
	// })

	ce := costexplorer.New(sess)

	costParams := &costexplorer.GetCostAndUsageInput{
		//Filter:      &costexplorer.Expression{},
		Granularity: aws.String("MONTHLY"),
		Metrics:     []*string{aws.String("UnblendedCost")},
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String("2019-09-01"),
			End:   aws.String("2019-11-01"),
		},
	}

	log.WithFields(log.Fields{
		"costParams": costParams,
	}).Debug("calling GetCostAndUsage")

	// Call to get detailed information on each instance
	cost, err := ce.GetCostAndUsage(costParams)
	if err != nil {
		return nil, err
	}

	sum := 0.0

	for _, element := range cost.ResultsByTime {
		//sum += strconv.Atoi(*element.Total["UnblendedCost"].Amount)
		amount := *element.Total["UnblendedCost"].Amount
		i, _ := strconv.ParseFloat(amount, 64)
		sum += i
		//fmt.Println("Total cost for last month", amount)
	}

	amountString := *cost.ResultsByTime[len(cost.ResultsByTime)-1].Total["UnblendedCost"].Amount
	amount, _ = strconv.ParseInt(amountString, 10, 64)

	// set check details
	outputReport.Payload.Check.Severity = checkcompleted.INFO
	outputReport.Payload.Check.Impact = int(amount)

	return &outputReport, nil

}
