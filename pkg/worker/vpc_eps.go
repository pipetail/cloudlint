package worker

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pipetail/cloudlint/pkg/awsregions"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func getVpcEndpointsWithinRegion(ec2client *ec2.EC2) []*ec2.VpcEndpoint {

	epsParams := &ec2.DescribeVpcEndpointsInput{
		Filters: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("attachment.status"),
				Values: []*string{
					aws.String("detached"),
				},
			},
		},
	}

	// Call to get detailed information on each volume
	eps, _ := ec2client.DescribeVpcEndpoints(epsParams)

	return eps.VpcEndpoints
}

// GetS3VpcEndpointsCount return a count of S3 VPC Endpoints
func GetS3VpcEndpointsCount(endpoints []*ec2.VpcEndpoint) int {

	count := 0

	if endpoints == nil {
		return count
	}

	if len(endpoints) == 0 {
		return count
	}

	for _, endpoint := range endpoints {

		if strings.HasSuffix(*endpoint.ServiceName, ".s3") {
			count++
		}

	}

	return count
}

// https://medium.com/nubego/how-to-save-money-with-aws-vpc-endpoints-9bac8ae1319c
func vpcendpoints(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	// externalID := event.Payload.AWSAuth.ExternalID
	// roleARN := event.Payload.AWSAuth.RoleARN

	s3endpointscount := 0

	log.WithFields(log.Fields{
		"event": event,
	}).Info("checking vpc_eps")

	//var countDisks int64 = 0
	var totalMonthlyPrice float64 = 0

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
		}).Debug("checking ebs_unused in aws region")

		ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})

		vpcendpoints := getVpcEndpointsWithinRegion(ec2Svc)

		// TODO: check if volumes.nextToken is nil
		// let's assume that each missing vpc endpoint costs 100$
		s3endpointscount += GetS3VpcEndpointsCount(vpcendpoints)
	}

	// TODO: make this relative to total spend
	severity := checkcompleted.INFO
	if s3endpointscount == 0 {
		severity = checkcompleted.WARNING
		totalMonthlyPrice = 100.0
	}

	// set check details
	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(totalMonthlyPrice)

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("vpc_eps_notused check finished")

	return &outputReport, nil

}
