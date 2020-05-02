package worker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/pipetail/cloudlint/pkg/awsregions"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func getDmsWithinRegion(client *databasemigrationservice.DatabaseMigrationService) []*databasemigrationservice.ReplicationInstance {

	params := &databasemigrationservice.DescribeReplicationInstancesInput{
		Filters:    []*databasemigrationservice.Filter{},
		Marker:     aws.String(""),
		MaxRecords: aws.Int64(20),
	}

	res, err := client.DescribeReplicationInstances(params)

	if err != nil {
		return nil
	}

	return res.ReplicationInstances
}

// GetDmsCount return a count of DMS Replication Instances
func GetDmsCount(instances []*databasemigrationservice.ReplicationInstance) (count int) {

	return len(instances)
}

// multi-AZ seems to be double
func getDmsComputePriceInRegion(region string, instance *databasemigrationservice.ReplicationInstance) (price float64) {

	// simplification for us-east-1
	instanceTypeMap := map[string]float64{
		"dms.t2.micro":   0.021,
		"dms.t2.small":   0.042,
		"dms.t2.medium":  0.084,
		"dms.t2.large":   0.168,
		"dms.c4.large":   0.197,
		"dms.c4.xlarge":  0.393,
		"dms.c4.2xlarge": 0.078,
		"dms.c4.4xlarge": 1.575,
		"dms.r4.large":   0.28,
		"dms.r4.xlarge":  0.55,
		"dms.r4.2xlarge": 1.11,
		"dms.r4.4xlarge": 2.22,
		"dms.r4.8xlarge": 4.43,
	}

	price = instanceTypeMap[*instance.ReplicationInstanceClass]

	// multi-AZ seems to be double
	if *instance.MultiAZ {
		price *= 2
	}

	// price per month
	return price * 24 * 30
}

// todo use region to get correct price
func getDmsStoragePriceInRegion(region string, instance *databasemigrationservice.ReplicationInstance) (price float64) {

	if instance == nil {
		return 0.0
	}

	// simplification for us-east-1
	price = float64(*instance.AllocatedStorage) * 0.115

	// multi-AZ seems to be double
	if *instance.MultiAZ {
		price *= 2
	}

	// price per month
	return price
}

func dmsUnused(event check.Event) (*checkcompleted.Event, error) {
	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	auth := event.Payload.AWSAuth
	var totalMonthlyPrice float64 = 0

	regions := awsregions.GetRegions()

	// see https://godoc.org/github.com/aws/aws-sdk-go/service/ec2#Region
	for _, region := range regions {
		log.WithFields(log.Fields{
			"awsRegion": region,
		}).Debug("checking dms_unused in aws region")

		//svc := databasemigrationservice.New(sess, &aws.Config{Credentials: creds, Region: aws.String(region)})
		svc := NewDMSClient(auth, region)

		instances := getDmsWithinRegion(svc)

		for _, instance := range instances {

			// Data transfer should be free
			totalMonthlyPrice += getDmsStoragePriceInRegion(region, instance) + getDmsComputePriceInRegion(region, instance)

			log.WithFields(log.Fields{
				"totalMonthlyPrice": totalMonthlyPrice,
			}).Debug("checking dms_unused in aws region: totalMonthlyPrice partial sum")
		}
	}

	// raise a warning if customer is using DMS
	severity := checkcompleted.INFO
	if totalMonthlyPrice > 0 {
		severity = checkcompleted.WARNING
	}

	// set check details
	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(totalMonthlyPrice)

	// you can also use DMS-free https://aws.amazon.com/dms/free-dms/

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("DMS Unused check finished")

	return &outputReport, nil
}
