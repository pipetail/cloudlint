package worker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/ec2"
	elb "github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/pipetail/cloudlint/pkg/checkreport"
)

// NewEC2Client constructs a new ec2 client with credentials and session
func NewEC2Client(auth checkreport.AwsAuth, region string) *ec2.EC2 {

	externalID := auth.ExternalID
	roleARN := auth.RoleARN

	sess := session.Must(session.NewSession())

	config := aws.NewConfig()

	if region != "" {
		config = config.WithRegion(region)
	}

	if (externalID != "") && (roleARN != "") {
		creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
			p.ExternalID = &externalID
		})

		config = config.WithCredentials(creds)
	}

	ec2Svc := ec2.New(sess, config)

	return ec2Svc
}

// NewELBClient constructs a new elb v2 client with credentials and session
func NewELBClient(auth checkreport.AwsAuth, region string) *elb.ELBV2 {

	externalID := auth.ExternalID
	roleARN := auth.RoleARN

	sess := session.Must(session.NewSession())

	config := aws.NewConfig()

	if region != "" {
		config = config.WithRegion(region)
	}

	if (externalID != "") && (roleARN != "") {
		creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
			p.ExternalID = &externalID
		})

		config = config.WithCredentials(creds)
	}

	elbSvc := elb.New(sess, config)

	return elbSvc
}

// NewDMSClient constructs a new dms client with credentials and session
func NewDMSClient(auth checkreport.AwsAuth, region string) *databasemigrationservice.DatabaseMigrationService {

	externalID := auth.ExternalID
	roleARN := auth.RoleARN

	sess := session.Must(session.NewSession())

	config := aws.NewConfig()

	if region != "" {
		config = config.WithRegion(region)
	}

	if (externalID != "") && (roleARN != "") {
		creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
			p.ExternalID = &externalID
		})

		config = config.WithCredentials(creds)
	}

	dmsSvc := databasemigrationservice.New(sess, config)

	return dmsSvc
}

// NewCEClient constructs a new costexplorer client with credentials and session
func NewCEClient(auth checkreport.AwsAuth) *costexplorer.CostExplorer {

	region := endpoints.UsEast1RegionID
	externalID := auth.ExternalID
	roleARN := auth.RoleARN

	sess := session.Must(session.NewSession())

	config := aws.NewConfig()

	if region != "" {
		config = config.WithRegion(region)
	}

	if (externalID != "") && (roleARN != "") {
		creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
			p.ExternalID = &externalID
		})

		config = config.WithCredentials(creds)
	}

	ceSvc := costexplorer.New(sess, config)

	return ceSvc
}
