package worker

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/pipetail/cloudlint/internal/pkg/check"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func paidSupport(event check.Event) (*checkcompleted.Event, error) {
	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	// externalID := event.Payload.AWSAuth.ExternalID
	// roleARN := event.Payload.AWSAuth.RoleARN

	impact := 0

	var err error

	// initialize severity
	severity := checkcompleted.INFO
	// authenticate to AWS
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsEast1RegionID),
	}))

	supportSvc := support.New(sess, &aws.Config{})
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
