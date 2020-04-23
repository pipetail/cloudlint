package worker

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
)

func checkIntegration() error {
	sess := session.Must(session.NewSession())
	costExplorerSvc := costexplorer.New(sess)
	_, err := checkIntegrationHandler(costExplorerSvc)
	if err != nil {
		return err
	}

	return nil
}

func checkIntegrationHandler(costExplorerScv costexploreriface.CostExplorerAPI) (*string, error) {

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

	log.Println(result)

	test := "bla"

	return &test, nil
}
