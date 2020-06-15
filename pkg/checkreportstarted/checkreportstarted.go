package checkreportstarted

import (
	"github.com/pipetail/cloudlint/pkg/check"
)

/*
{
    "name": "CheckReportStarted",
    "payload": {
        "reportId": "<Uuid>",
        "checks": [
            {
                "id": "<Uuid>",
                "type": "<String>",
                "group": "<String>",
                "name": "<String>",
                "description": "<String>"
            }
        ]
    }
}
*/

// Event struct for CheckReportStarted event
type Event struct {
	Name    string  `json:"name"`
	Payload Payload `json:"payload"`
}

// Payload of CheckReportStarted event
type Payload struct {
	ReportID string        `json:"reportId"`
	Checks   []check.Check `json:"checks"`
}

// New - constructor for CheckReportStarted event with default values
func New(reportID string) Event {
	return Event{
		Name: "CheckReportStarted",
		Payload: Payload{
			Checks:   check.GetChecks(),
			ReportID: reportID,
		},
	}
}

// WithFilter - apply filter on checks
func (e Event) WithFilter(filterCheck []string) Event {

	// no filter specified, return the original event
	if len(filterCheck) == 0 {
		return e
	}

	newEvent := New(e.Payload.ReportID)

	newEvent.Payload.Checks = nil

	for _, check := range e.Payload.Checks {

		keep := func() bool {
			for _, filter := range filterCheck {
				if filter == check.Type {
					return true
				}
			}
			return false
		}

		if keep() {
			newEvent.Payload.Checks = append(newEvent.Payload.Checks, check)
		}

	}

	return newEvent
}
