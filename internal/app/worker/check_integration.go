package worker

import (
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

func checkIntegrationHandler(costExplorerScv costexploreriface.CostExplorerAPI) (string, error) {
	return "bla", nil
}
