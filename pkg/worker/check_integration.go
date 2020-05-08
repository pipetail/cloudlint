package worker

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"

	"github.com/pipetail/cloudlint/internal/utils"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
)

func checkIntegration(event check.Event) (*checkcompleted.Event, error) {
	// prepare output report
	outputReport := checkcompleted.New(event.Payload.CheckID)
	outputReport.Name = "checkIntegration"

	sess := session.Must(session.NewSession())
	costExplorerSvc := costexplorer.New(sess)
	err := checkIntegrationHandler(costExplorerSvc, &outputReport)
	return &outputReport, err
}

func checkIntegrationHandler(costExplorerScv costexploreriface.CostExplorerAPI, event *checkcompleted.Event) error {

	// create dummy query to costexplorer
	costParams := &costexplorer.GetCostAndUsageInput{
		Granularity: aws.String("MONTHLY"),
		Metrics:     []*string{aws.String("UnblendedCost")},
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String(utils.GetLastBillingPeriodStart()),
			End:   aws.String(utils.GetLastBillingPeriodEnd()),
		},
	}

	result, err := costExplorerScv.GetCostAndUsage(costParams)
	if err != nil {
		return fmt.Errorf("GetCostAndUsage: %s", err)
	}

	sum := 0.0
	for _, element := range result.ResultsByTime {
		amount := *element.Total["UnblendedCost"].Amount
		i, _ := strconv.ParseFloat(amount, 64)
		sum += i
	}

	// adjust event if needed here
	// event.Payload.Check...

	return nil
}
