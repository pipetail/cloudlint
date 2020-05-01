package checkawsintegrationcompleted

/*
{
    "name": "CheckAwsIntegrationCompleted",
    "payload": {
        "requestId": "<Uuid>",
        "status": "<Int>"
    }
}

*/

// Status - SUCCESS or FAIL
type Status int

// Status iota
const (
	SUCCESS Status = iota // SUCCESS == 0
	FAIL                  // FAIL == 1
)

// Event for CheckAwsIntegrationCompleted
type Event struct {
	Name    string  `json:"name"`
	Payload Payload `json:"payload"`
}

// Payload of the event
type Payload struct {
	Status    Status `json:"status"`
	RequestID string `json:"requestId"`
}

// New - constructor for creating CheckAwsIntegrationCompleted event
func New(requestID string, status Status) Event {
	return Event{
		Name: "CheckAwsIntegrationCompleted",
		Payload: Payload{
			RequestID: requestID,
			Status:    status,
		},
	}
}
