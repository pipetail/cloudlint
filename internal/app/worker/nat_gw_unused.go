package worker

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/pipetail/cloudlint/internal/pkg/awsregions"
	"github.com/pipetail/cloudlint/internal/pkg/check"
	"github.com/pipetail/cloudlint/internal/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func getNatGatewaysWithinRegion(ec2client *ec2.EC2) []*ec2.NatGateway {

	// TODO: now we are only checking ALL THE EXISTING nat gateways that are available

	// we should do a more complex check where NAT Gateway has 0 traffic flowing through
	// or sits in a subnet with associated private subnet that is empty

	// or that private subnet has no 0.0.0.0/0 route to NAT GW

	// you can check some metrics from CloudWatch too
	// https://docs.aws.amazon.com/vpc/latest/userguide/vpc-nat-gateway-cloudwatch.html
	epsParams := &ec2.DescribeNatGatewaysInput{
		Filter: []*ec2.Filter{
			&ec2.Filter{
				Name: aws.String("status"),
				Values: []*string{
					aws.String("available"),
				},
			},
		},
	}

	// Call to get detailed information on each volume
	eps, _ := ec2client.DescribeNatGateways(epsParams)

	return eps.NatGateways
}

// GetNatGatewaysCount return a count of NAT Gateways
func GetNatGatewaysCount(natgateways []*ec2.NatGateway) (count int) {

	return len(natgateways)
}

// https://aws.amazon.com/premiumsupport/knowledge-center/vpc-reduce-nat-gateway-transfer-costs/

func getNatGatewayPriceInRegion(region string) (price float64) {

	priceMap := map[string]float64{
		"us-east-1":      0.045,
		"us-east-2":      0.045,
		"us-west-1":      0.048,
		"us-west-2":      0.045,
		"ap-east-1":      0.065,
		"ap-south-1":     0.056,
		"ap-northeast-3": 0.062,
		"ap-northeast-2": 0.059,
		"ap-southeast-1": 0.059,
		"ap-southeast-2": 0.059,
		"ap-northeast-1": 0.062,
		"ca-central-1":   0.05,
		//"cn-north-1":
		//"cn-northwest-1":
		"eu-central-1":  0.052,
		"eu-west-1":     0.048,
		"eu-west-2":     0.05,
		"eu-west-3":     0.05,
		"eu-north-1":    0.046,
		"me-south-1":    0.0528,
		"sa-east-1":     0.093,
		"us-gov-east-1": 0.054,
		"us-gov-west-1": 0.054,
	}

	return priceMap[region]

}

func natgwunused(event check.Event) (*checkcompleted.Event, error) {

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	// externalID := event.Payload.AWSAuth.ExternalID
	// roleARN := event.Payload.AWSAuth.RoleARN

	natgwscount := 0

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
		}).Info("checking ebs_unused in aws region")

		ec2Svc := ec2.New(sess, &aws.Config{Region: aws.String(region)})

		natgateways := getNatGatewaysWithinRegion(ec2Svc)

		// TODO: check if .nextToken is nil
		natgwscount += GetNatGatewaysCount(natgateways)

		// count the price
		// those are unused nat gateways (==no traffic flowing through it, so that should be free)
		// you only pay $/hour then
		totalMonthlyPrice += float64(natgwscount) * getNatGatewayPriceInRegion(region) * (24 * 30)
	}

	// TODO: make this relative to total spend
	severity := checkcompleted.INFO
	if natgwscount != 0 {
		severity = checkcompleted.WARNING
	}

	// set check details
	outputReport.Payload.Check.Severity = severity
	outputReport.Payload.Check.Impact = int(totalMonthlyPrice)

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("nat_gw_unused check finished")

	return &outputReport, nil

}
