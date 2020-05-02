package worker

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pipetail/cloudlint/pkg/awsregions"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func getAddressesWithingRegion(ec2client *ec2.EC2) []*ec2.Address {

	params := &ec2.DescribeAddressesInput{
		Filters: []*ec2.Filter{},
	}

	res, _ := ec2client.DescribeAddresses(params)

	return res.Addresses
}

func filterUnassociatedAddresses(input []*ec2.Address) []*ec2.Address {

	// we filter them out by creating new slice
	unassociated := make([]*ec2.Address, 0)

	for _, address := range input {

		// and adding only the addresses that are not associated
		if address.AssociationId == nil {
			unassociated = append(unassociated, address)
		}
	}

	return unassociated
}

// GetAddressesCount return a count unassociated EIPs
func GetAddressesCount(addresses []*ec2.Address) (count int) {

	unassociated := filterUnassociatedAddresses(addresses)

	return len(unassociated)
}

// https://aws.amazon.com/premiumsupport/knowledge-center/vpc-reduce-nat-gateway-transfer-costs/

func getAddressPriceInRegion(region string) (price float64) {

	priceMap := map[string]float64{
		"us-east-1": 0.005,
	}

	// simplification, TODO: add all regions
	return priceMap["us-east-1"]

}

func eipunused(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	auth := event.Payload.AWSAuth

	eipcount := 0

	log.WithFields(log.Fields{
		"event": event,
	}).Info("checking eip_unused")

	//var countDisks int64 = 0
	var totalMonthlyPrice float64 = 0

	regions := awsregions.GetRegions()

	// see https://godoc.org/github.com/aws/aws-sdk-go/service/ec2#Region
	for _, region := range regions {

		log.WithFields(log.Fields{
			"awsRegion": region,
		}).Debug("checking eip_unused in aws region")

		ec2Svc := NewEC2Client(auth, region)

		addresses := getAddressesWithingRegion(ec2Svc)

		// TODO: check if .nextToken is nil
		eipCountWithinRegion := GetAddressesCount(addresses)
		eipcount += GetAddressesCount(addresses)

		// count the price
		totalMonthlyPrice += float64(eipCountWithinRegion) * getAddressPriceInRegion(region) * (24 * 30)
	}

	// TODO: make this relative to total spend
	severity := checkcompleted.INFO
	if eipcount != 0 {
		severity = checkcompleted.WARNING
	}

	// set check details
	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(totalMonthlyPrice)

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("eip_unused check finished")

	return &outputReport, nil

}
