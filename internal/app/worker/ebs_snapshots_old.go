package worker

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/pipetail/cloudlint/internal/pkg/check"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func ebsSnapshotsOld(event check.Event) (*checkcompleted.Event, error) {
	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	log.WithFields(log.Fields{
		"report": outputReport,
	}).Info("starting ebs_snapshots_old")

	// externalID := event.Payload.AWSAuth.ExternalID
	// roleARN := event.Payload.AWSAuth.RoleARN
	var impact float64

	// log.WithFields(log.Fields{
	// 	"roleARN": roleARN,
	// }).Info("checking with roleARN")

	// authenticate to AWS
	sess := session.Must(session.NewSession())
	// creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
	// 	p.ExternalID = &externalID
	// })

	// TODO: add region support

	// create AWS SDK clients
	ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String("us-east-1")})
	stsSvc := sts.New(sess)

	// get account id
	getCallerIdenityInput := sts.GetCallerIdentityInput{}
	identity, err := stsSvc.GetCallerIdentity(&getCallerIdenityInput)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
		}).Error("could not obtain sts info")
		return nil, fmt.Errorf("could not obtain sts info: %s", err)
	}

	// list all snapshots for the given owners
	describeSnapshotInput := ec2.DescribeSnapshotsInput{
		OwnerIds: []*string{identity.Account},
	}
	snapshots, err := ec2Svc.DescribeSnapshots(&describeSnapshotInput)
	if err != nil {
		return nil, fmt.Errorf("could not obtain snapshots: %s", err)
	}

	totalSize := int64(0)
	for _, snapshot := range snapshots.Snapshots {
		log.WithFields(log.Fields{
			"SnapshotId": *snapshot.SnapshotId,
		}).Info("found snapshot")

		// older than 24 hours?
		if time.Now().Sub(*snapshot.StartTime).Hours() > 90*24 {
			totalSize = totalSize + *snapshot.VolumeSize
		}
	}

	// https://aws.amazon.com/ebs/pricing/ Frankfurt
	impact = float64(totalSize) * float64(0.054)

	// set severity based on total size of snapshots
	severity := checkcompleted.INFO
	if impact > 0 {
		severity = checkcompleted.WARNING
	}

	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(impact)

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("EBS snapshots old check finished")

	return &outputReport, nil
}
