package checkreport

/*
{
    "name": "CheckReport",
    "payload": {
        "awsAuth": {
            "roleArn": "<String>",
            "externalId": "<String>"
        }
        "reportId": "<Uuid>",
    }
}
*/

// Event for CheckReport command
type Event struct {
	Name    string  `json:"name"`
	Payload Payload `json:"payload"`
}

// Payload of CheckReport command
type Payload struct {
	AWSAuth  AwsAuth `json:"awsAuth"`
	ReportID string  `json:"reportId"`
}

// AwsAuth struct to hold IAM auth info
type AwsAuth struct {
	RoleARN    string `json:"roleArn"`
	ExternalID string `json:"externalId"`
}

// New - constructor for creating CheckReport event
func New(reportID string) Event {
	return Event{
		Name: "CheckReport",
		Payload: Payload{
			ReportID: reportID,
		},
	}
}
