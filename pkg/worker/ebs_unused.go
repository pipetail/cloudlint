package worker

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
	"github.com/aws/aws-sdk-go/service/pricing/pricingiface"
	"github.com/pipetail/cloudlint/pkg/awspricing"
	"github.com/pipetail/cloudlint/pkg/awsregions"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	ins "github.com/pipetail/cloudlint/pkg/inspection"
	log "github.com/sirupsen/logrus"
)

// Parameters for AWS session used for dependency injection
type Parameters struct {
	session     *session.Session
	credentials *credentials.Credentials
	region      *string
}

// GetVolumesPrice sums the final price for all the volumes
func GetVolumesPrice(volumes []*ec2.Volume, client pricingiface.PricingAPI, region string) float64 {

	var totalSize int64 = 0
	var totalMonthlyPrice float64 = 0

	for _, volume := range volumes {

		totalSize += *volume.Size

		totalMonthlyPrice += priceOfVolume(volume, client, region)
	}

	return totalMonthlyPrice
}

func priceOfVolume(volume *ec2.Volume, client pricingiface.PricingAPI, region string) float64 {
	return float64(*volume.Size) * awspricing.GetPriceOfVolume(client, *volume.VolumeType, region)
}

func getVolumesWithinRegion(client ec2iface.EC2API) []*ec2.Volume {

	// we need all volumes, we will filter them later on
	volumeParams := &ec2.DescribeVolumesInput{}

	// Call to get detailed information on each volume
	res, _ := client.DescribeVolumes(volumeParams)

	return res.Volumes
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

	pricingClient := NewPricingClient(auth)

	//var countDisks int64 = 0
	var totalMonthlyPrice float64 = 0

	regions := awsregions.GetRegions()
	details := make([]checkcompleted.Detail, 0)

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
		if ins.CheckDetail() {
			details = append(details, readDetail(detachedVolumes, pricingClient, region)...)
		}
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
	outputReport.Payload.Check.Details = details

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("EBS unused check finished")

	return &outputReport, nil
}

func readDetail(volumes []*ec2.Volume, client pricingiface.PricingAPI, region string) []checkcompleted.Detail {
	details := make([]checkcompleted.Detail, 0, len(volumes))

	for _, volume := range volumes {
		details = append(details, checkcompleted.Detail{
			Region: region,
			ID:     *volume.VolumeId,
			Size:   fmt.Sprintf("%d", *volume.Size),
			Cost:   fmt.Sprintf("%.2f", priceOfVolume(volume, client, region)),
			Tags:   mapTags(volume.Tags),
		})
	}
	return details
}

func mapTags(tags []*ec2.Tag) []string {
	details := make([]string, 0, len(tags))
	for _, tag := range tags {
		if *tag.Key == "Name" {
			details = append(details, fmt.Sprintf("%s", *tag.Value))
		}
	}
	return details
}
