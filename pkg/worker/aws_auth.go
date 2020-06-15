package worker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/ec2"
	elb "github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/aws/aws-sdk-go/service/pricing"
	"github.com/aws/aws-sdk-go/service/support"
	"github.com/pipetail/cloudlint/pkg/checkreport"
	log "github.com/sirupsen/logrus"
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

	log.WithFields(log.Fields{
		"ec2Svc": ec2Svc,
	}).Debug("ec2 client init")

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

	log.WithFields(log.Fields{
		"elbSvc": elbSvc,
	}).Debug("elb client init")

	return elbSvc
}

// NewCWClient constructs a new cloudwatch client with credentials and session
func NewCWClient(auth checkreport.AwsAuth, region string) *cloudwatch.CloudWatch {

	externalID := auth.ExternalID
	roleARN := auth.RoleARN

	sess := session.Must(session.NewSession())

	config := aws.NewConfig().WithCredentialsChainVerboseErrors(true)

	if region != "" {
		config = config.WithRegion(region)
	}

	if (externalID != "") && (roleARN != "") {
		creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
			p.ExternalID = &externalID
		})

		config = config.WithCredentials(creds)
	}

	svc := cloudwatch.New(sess, config)

	log.WithFields(log.Fields{
		"cloudwatch": svc,
	}).Debug("cloudwatch client init")

	return svc
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

	log.WithFields(log.Fields{
		"dmsSvc": dmsSvc,
	}).Debug("dms client init")

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

	log.WithFields(log.Fields{
		"ceSvc": ceSvc,
	}).Debug("ce client init")

	return ceSvc
}

// NewSupportClient constructs a new support client with credentials and session
func NewSupportClient(auth checkreport.AwsAuth) *support.Support {

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

	svc := support.New(sess, config)

	log.WithFields(log.Fields{
		"supportSvc": svc,
	}).Debug("support client init")

	return svc
}

// NewPricingClient constructs a new pricing client with credentials and session
func NewPricingClient(auth checkreport.AwsAuth) *pricing.Pricing {

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

	svc := pricing.New(sess, config)

	log.WithFields(log.Fields{
		"supportSvc": svc,
	}).Debug("support client init")

	return svc
}
