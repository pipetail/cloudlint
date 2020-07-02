package worker

import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
    "github.com/aws/aws-sdk-go/service/pricing"
    "github.com/pipetail/cloudlint/internal/utils"
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

        totalMonthlyPrice += float64(*volume.Size) * getPriceOfValue(client, *volume.VolumeType, region)
	}

	return totalMonthlyPrice
}

func getPriceOfValue(client *pricing.Pricing, volumeType string, region string) float64 {

    input := pricing.GetProductsInput{
        Filters: []*pricing.Filter{
            {
                Field: aws.String("ServiceCode"),
                Type:  aws.String("TERM_MATCH"),
                Value: aws.String("AmazonEC2"),
            },
            {
                Field: aws.String("Location"),
                Type:  aws.String("TERM_MATCH"),
                Value: aws.String(utils.GetLocationForRegion(region)),
            },
            {
                Field: aws.String("volumeType"),
                Type:  aws.String("TERM_MATCH"),
                Value: aws.String(volumeType),
            },
        },
        FormatVersion: aws.String("aws_v1"),
        MaxResults:    aws.Int64(1),
    }

    // this is a workaround for a bug: https://github.com/aws/aws-sdk-go/issues/3323
    input.SetServiceCode("AmazonEC2")

    log.WithFields(log.Fields{
        "input": input,
    }).Info("getPriceOfValue")

    resp, err := client.GetProducts(&input)

    if err != nil {
        log.WithFields(log.Fields{
            "err": err,
        }).Error("checking getPriceOfValue")
        return 0
    }

    price := extractPrice(resp)

    pricePerMonth := getPricePerMonth(price)

    return pricePerMonth
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
