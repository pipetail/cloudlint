package checkcompleted

import "fmt"

/*
{
    "name": "CheckCompleted",
    "payload": {
        "check": {
            "id": "<Uuid>",
            "severity": "<Int>",
            "impact": "<Int>"
        }
    }
}

*/

// Severity INFO, WARNING, ERROR
type Severity int

// Severity INFO, WARNING, ERROR
const (
	INFO    Severity = iota // INFO == 0
	WARNING                 // WARNING == 1
	ERROR                   // ERROR == 2
)

func (e Severity) String() string {
	switch e {
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	default:
		return fmt.Sprintf("%d", int(e))
	}
}

// Event checkcompleted event
type Event struct {
	Name    string  `json:"name"`
	Payload Payload `json:"payload"`
}

// Payload of the event
type Payload struct {
	Check Check `json:"check"`
}

// Check struct
type Check struct {
	ID       string   `json:"id"`
	Severity Severity `json:"severity"`
	Impact   int      `json:"impact"`
}

// TODO: delete this
// New constructs a CheckCompleted event
func New(id string) Event {
	return Event{
		Name: "CheckCompleted",
		Payload: Payload{
			Check: Check{
				ID:       id,
				Severity: INFO,
				Impact:   0,
			},
		},
	}
}
