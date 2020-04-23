package worker

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/pipetail/cloudlint/internal/pkg/check"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"

	"strconv"

	"github.com/aws/aws-sdk-go/service/costexplorer"
)

func billingInfo(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	externalID := event.Payload.AWSAuth.ExternalID
	roleARN := event.Payload.AWSAuth.RoleARN
	var amount int64

	log.WithFields(log.Fields{
		"roleARN": roleARN,
	}).Debug("checking with roleARN")

	if roleARN != "arn:aws:iam::680177765279:role/awsdemo" {

		//Create new Cost Explorer client
		// authenticate to AWS
		sess := session.Must(session.NewSession())
		creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
			p.ExternalID = &externalID
		})

		ce := costexplorer.New(sess, &aws.Config{Credentials: creds})

		costParams := &costexplorer.GetCostAndUsageInput{
			//Filter:      &costexplorer.Expression{},
			Granularity: aws.String("MONTHLY"),
			Metrics:     []*string{aws.String("UnblendedCost")},
			TimePeriod: &costexplorer.DateInterval{
				Start: aws.String("2019-09-01"),
				End:   aws.String("2019-11-01"),
			},
		}
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
			fmt.Println("Total cost for last month", amount)
		}

		//fmt.Println("Total cost for three consecutive months: ", totalSize, "GiB")
		amountString := *cost.ResultsByTime[len(cost.ResultsByTime)-1].Total["UnblendedCost"].Amount
		amount, _ = strconv.ParseInt(amountString, 10, 64)
		//currency := *cost.ResultsByTime[len(cost.ResultsByTime)-1].Total["UnblendedCost"].Unit
		//fmt.Println("Total cost for last month", amount, currency)
		//fmt.Println("Total cost for last three months", sum, currency)

	} else {

		amount = 15874

	}

	log.WithFields(log.Fields{
		"Amount":  amount,
		"CheckID": event.Payload.CheckID,
	}).Debug("Checked Unblended cost for all services in cost explorer")

	// set check details
	outputReport.Payload.Check.Severity = checkcompleted.INFO
	outputReport.Payload.Check.Impact = int(amount)

	return &outputReport, nil

}
