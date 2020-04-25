package worker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
)

func (client *CostExplorerMockClient) GetCostAndUsage(input *costexplorer.GetCostAndUsageInput) (*costexplorer.GetCostAndUsageOutput, error) {

	payload := `{
		"ResultsByTime": [
			{
				"TimePeriod": {
					"Start": "2019-09-01",
					"End": "2019-10-01"
				},
				"Total": {
					"UnblendedCost": {
						"Amount": "0.0000030803",
						"Unit": "USD"
					}
				},
				"Groups": [],
				"Estimated": false
			},
			{
				"TimePeriod": {
					"Start": "2019-10-01",
					"End": "2019-11-01"
				},
				"Total": {
					"UnblendedCost": {
						"Amount": "7782.9981148686",
						"Unit": "USD"
					}
				},
				"Groups": [],
				"Estimated": false
			}
		]
	}`

	output := costexplorer.GetCostAndUsageOutput{}
	err := json.Unmarshal([]byte(payload), &output)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %s", err)
	}

	return &output, nil
}

func TestCheckIntegrationHandler(t *testing.T) {
	mockClient := CostExplorerMockClient{}
	outputReport := checkcompleted.New("dummyID")
	err := checkIntegrationHandler(&mockClient, &outputReport)

	if checkIntegrationHandler(&mockClient, &outputReport) != nil {
		t.Errorf("could not verify integration status: %s", err)
	}
}
