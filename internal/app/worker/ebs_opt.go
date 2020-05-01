package worker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pipetail/cloudlint/pkg/awsregions"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

// InstanceAtribute struct
type InstanceAtribute struct {
	InstanceID   string `json:"instanceId"`
	EbsOptimized bool   `json:"ebsOptimized"`
}

func checkEbsNotOptimized(attributes []*InstanceAtribute) bool {
	for _, attribute := range attributes {
		if !attribute.EbsOptimized {
			return true
		}
	}
	return false
}

func getEbsOptimizedForInstanceAtributes(ec2client *ec2.EC2) []*InstanceAtribute {

	// all atributes of instances from attachments
	attributes := make([]*InstanceAtribute, 0)

	volumes := getVolumesWithinRegion(ec2client)

	for _, volume := range volumes {

		// for attachments find if any
		if len(volume.Attachments) > 0 {
			for _, attachment := range volume.Attachments {
				id := attachment.InstanceId

				input := &ec2.DescribeInstanceAttributeInput{
					Attribute:  aws.String("ebsOptimized"),
					InstanceId: id,
				}

				result, _ := ec2client.DescribeInstanceAttribute(input)

				attributes = append(attributes, &InstanceAtribute{
					InstanceID:   *id,
					EbsOptimized: *result.EbsOptimized.Value,
				})

			}
		}
	}

	return attributes
}

func ebsopt(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	// externalID := event.Payload.AWSAuth.ExternalID
	// roleARN := event.Payload.AWSAuth.RoleARN

	severity := checkcompleted.INFO

	// authenticate to AWS
	sess := session.Must(session.NewSession())
	// creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
	// 	p.ExternalID = &externalID
	// })

	regions := awsregions.GetRegions()

	// see https://godoc.org/github.com/aws/aws-sdk-go/service/ec2#Region
	for _, region := range regions {

		log.WithFields(log.Fields{
			"awsRegion": region,
		}).Debug("checking ebs_opt in aws region")

		ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})

		attributes := getEbsOptimizedForInstanceAtributes(ec2Svc)

		if checkEbsNotOptimized(attributes) {
			severity = checkcompleted.WARNING
		}
	}

	// set check details
	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(0)

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("EBS optimized check finished")

	return &outputReport, nil

}
