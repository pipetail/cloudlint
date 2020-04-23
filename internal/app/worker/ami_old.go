package worker

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/pipetail/cloudlint/internal/pkg/awsregions"
	"github.com/pipetail/cloudlint/internal/pkg/check"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func countImagesUsedSize(images []*ec2.Image, threshold int64) (int64, error) {
	count := int64(0)
	for _, image := range images {
		for _, mapping := range image.BlockDeviceMappings {
			if mapping.Ebs != nil {
				// add more device types
				created, err := time.Parse(time.RFC3339, *image.CreationDate)
				if err != nil {
					return 0, err
				}
				if time.Now().Sub(created).Hours() > float64(threshold) {
					count = count + *mapping.Ebs.VolumeSize
				}
			}
		}
	}

	return count, nil
}

func getImagePriceInRegion(region string) (price float64) {

	// monthly price for 1 GiB
	// images are stored as snapshots (in s3 or similar)
	// https://aws.amazon.com/ebs/pricing/

	priceMap := map[string]float64{
		"us-east-1":      0.05,
		"us-east-2":      0.05,
		"us-west-1":      0.055,
		"us-west-2":      0.05,
		"ap-east-1":      0.055,
		"ap-south-1":     0.05,
		"ap-northeast-3": 0.05,
		"ap-northeast-2": 0.05,
		"ap-southeast-1": 0.05,
		"ap-southeast-2": 0.055,
		"ap-northeast-1": 0.05,
		"ca-central-1":   0.055,
		"eu-central-1":   0.054,
		"eu-west-1":      0.05,
		"eu-west-2":      0.053,
		"eu-west-3":      0.053,
		"eu-north-1":     0.0475,
		"me-south-1":     0.055,
		"sa-east-1":      0.068,
		"us-gov-east-1":  0.066,
		"us-gov-west-1":  0.066,
	}

	return priceMap[region]
}

func getImagesWithinRegion(ec2client *ec2.EC2, stsclient *sts.STS) ([]*ec2.Image, error) {
	// get account id
	getCallerIdenityInput := sts.GetCallerIdentityInput{}
	identity, err := stsclient.GetCallerIdentity(&getCallerIdenityInput)
	if err != nil {
		return nil, fmt.Errorf("could not obtain sts info: %s", err)
	}

	// we can use the same input for all regions
	describeImagesInput := ec2.DescribeImagesInput{
		Owners: []*string{identity.Account},
	}

	// get images
	result, err := ec2client.DescribeImages(&describeImagesInput)
	if err != nil {
		return nil, fmt.Errorf("could not obtain images: %s", err)
	}

	return result.Images, nil
}

func amiOld(event check.Event) (*checkcompleted.Event, error) {
	outputReport := checkcompleted.New(event.Payload.CheckID)
	// externalID := event.Payload.AWSAuth.ExternalID
	// roleARN := event.Payload.AWSAuth.RoleARN

	sess := session.Must(session.NewSession())
	// creds := stscreds.NewCredentials(sess, roleARN, func(p *stscreds.AssumeRoleProvider) {
	// 	p.ExternalID = &externalID
	// })

	// prepare STS client for all regions
	stsSvc := sts.New(sess)

	// prepare counter for all regions
	impactTotal := int64(0)

	// check all regions
	regions := awsregions.GetRegions()
	for _, region := range regions {
		log.WithFields(log.Fields{
			"awsRegion": region,
		}).Debug("checking ami_old in aws region")

		// create svc for the given region
		ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})
		images, err := getImagesWithinRegion(ec2Svc, stsSvc)
		if err != nil {
			return nil, fmt.Errorf("could not obtain images: %s", err)
		}

		// count space used by images in the current region
		countRegional, err := countImagesUsedSize(images, 90*24) // 90 days
		if err != nil {
			return nil, fmt.Errorf("could not obtain count of used EBS space for %s: %s", region, err)
		}

		// get total price for the current region (snapshots have monthly pricing / GiB)
		impactRegional := getImagePriceInRegion(region) * float64(countRegional)
		impactTotal = impactTotal + int64(impactRegional)
	}

	// set report status
	severity := checkcompleted.INFO
	if impactTotal != 0 {
		severity = checkcompleted.WARNING
	}

	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(impactTotal)

	return &outputReport, nil
}
