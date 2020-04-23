package worker

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
)

type CostExplorerMockClient struct {
	costexploreriface.CostExplorerAPI
}

func TestCheckIntegrationHandler(t *testing.T) {
	mockClient := CostExplorerMockClient{}
	_, err := checkIntegrationHandler(&mockClient)

	if err != nil {
		t.Errorf("bla: %s", err)
	}
}
