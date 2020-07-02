package worker

import (
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/pricing"
    "github.com/pipetail/cloudlint/pkg/awspricing"
    "github.com/pipetail/cloudlint/pkg/awsregions"
    "github.com/pipetail/cloudlint/pkg/check"
    "github.com/pipetail/cloudlint/pkg/checkcompleted"
    log "github.com/sirupsen/logrus"
)

// Parameters for AWS session used for dependency injection
type Parameters struct {
	session     *session.Session
	credentials *credentials.Credentials
	region      *string
}

// GetVolumesPrice sums the final price for all the volumes
func GetVolumesPrice(volumes []*ec2.Volume, client *pricing.Pricing, region string) float64 {

	var totalSize int64 = 0
	var totalMonthlyPrice float64 = 0

	for _, volume := range volumes {

		totalSize += *volume.Size
		//countDisks++

        totalMonthlyPrice += float64(*volume.Size) * awspricing.GetPriceOfValue(client, *volume.VolumeType, region)
	}

	return totalMonthlyPrice
}

func getVolumesWithinRegion(ec2client *ec2.EC2) []*ec2.Volume {

	volumeParams := &ec2.DescribeVolumesInput{}

	// Call to get detailed information on each volume
	volumes, _ := ec2client.DescribeVolumes(volumeParams)

	return volumes.Volumes
}

func filterDetachedVolumes(input []*ec2.Volume) []*ec2.Volume {

	// we filter them out by creating new slice
	attached := make([]*ec2.Volume, 0)

	for _, volume := range input {

		// and adding only the volumes that are detached
		if len(volume.Attachments) == 0 {
			attached = append(attached, volume)
		}
	}

	return attached
}

func ebsunused(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	auth := event.Payload.AWSAuth

	//var countDisks int64 = 0
	var totalMonthlyPrice float64 = 0

	regions := awsregions.GetRegions()

    pricingClient := NewPricingClient(auth)

	// see https://godoc.org/github.com/aws/aws-sdk-go/service/ec2#Region
	for _, region := range regions {

		log.WithFields(log.Fields{
			"awsRegion": region,
		}).Debug("checking ebs_unused in aws region")

		ec2Svc := NewEC2Client(auth, region)

		volumes := getVolumesWithinRegion(ec2Svc)

		detachedVolumes := filterDetachedVolumes(volumes)

		// TODO: check if volumes.nextToken is nil
		totalMonthlyPrice += GetVolumesPrice(detachedVolumes, pricingClient, region)
        awspricing.GetPriceOfValue(pricingClient, "Provisioned IOPS", region)
    }

	// TODO: make this relative to total spend
	severity := checkcompleted.INFO
	if totalMonthlyPrice > 20 {
		severity = checkcompleted.WARNING
	}
	if totalMonthlyPrice > 100 {
		severity = checkcompleted.ERROR
	}

	// set check details
	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(totalMonthlyPrice)

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("EBS unused check finished")

	return &outputReport, nil
}
