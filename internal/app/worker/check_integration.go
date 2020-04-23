package worker

import (
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
)

func checkIntegration() (*checkcompleted.Event, error) {
	sess := session.Must(session.NewSession())
	costExplorerSvc := costexplorer.New(sess)
	return checkIntegrationHandler(costExplorerSvc)
}

func checkIntegrationHandler(costExplorerScv costexploreriface.CostExplorerAPI) (*checkcompleted.Event, error) {

	costParams := &costexplorer.GetCostAndUsageInput{
		Granularity: aws.String("MONTHLY"),
		Metrics:     []*string{aws.String("UnblendedCost")},
		TimePeriod: &costexplorer.DateInterval{
			Start: aws.String("2019-09-01"),
			End:   aws.String("2019-11-01"),
		},
	}

	result, err := costExplorerScv.GetCostAndUsage(costParams)
	if err != nil {
		return nil, fmt.Errorf("GetCostAndUsage: %s", err)
	}

	sum := 0.0
	for _, element := range result.ResultsByTime {
		amount := *element.Total["UnblendedCost"].Amount
		i, _ := strconv.ParseFloat(amount, 64)
		sum += i
	}

	output := checkcompleted.Event{
		Name: "checkIntegration",
	}

	return &output, nil
}
