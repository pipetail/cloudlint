package checkawsintegration

import "github.com/pipetail/cloudlint/internal/pkg/checkreport"

/*
{
    "name": "CheckAwsIntegration",
    "payload": {
        "requestId": "<Uuid>",
        "awsAuth": {
            "roleArn": "<String>",
            "externalId": "<String>"
        }
    }
}

*/

// Event of CheckAwsIntegration
type Event struct {
	Name    string  `json:"name"`
	Payload Payload `json:"payload"`
}

// Payload of CheckAwsIntegration event
type Payload struct {
	AWSAuth   checkreport.AwsAuth `json:"awsAuth"`
	RequestID string              `json:"requestId"`
}
