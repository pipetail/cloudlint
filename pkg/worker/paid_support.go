package worker

import (
	"strings"

	"github.com/aws/aws-sdk-go/service/support"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func paidSupport(event check.Event) (*checkcompleted.Event, error) {
	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	auth := event.Payload.AWSAuth

	impact := 0

	var err error

	// initialize severity
	severity := checkcompleted.INFO

	supportSvc := NewSupportClient(auth)
	_, err = supportSvc.DescribeServices(&support.DescribeServicesInput{})

	// reverse logic, we need to get certain error message
	// so we know that support is not enabled
	if err != nil {
		if strings.Contains(err.Error(), "SubscriptionRequiredException") {
			severity = checkcompleted.INFO
		} else {
			// this is the real error, fail the function
			return nil, err
		}
	} else {
		// paid support is enabled
		severity = checkcompleted.WARNING
	}

	// set check details
	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = impact

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Debug("Paid support check finished")

	return &outputReport, nil
}
