package worker

import (
	"fmt"

	elb "github.com/aws/aws-sdk-go/service/elbv2"
	"github.com/pipetail/cloudlint/pkg/awsregions"
	"github.com/pipetail/cloudlint/pkg/check"
	"github.com/pipetail/cloudlint/pkg/checkcompleted"
	log "github.com/sirupsen/logrus"
)

func elbUnused(event check.Event) (*checkcompleted.Event, error) {

	log.WithFields(log.Fields{
		"LoadBalancerArn": "",
	}).Info("starting elb_unused check")

	// prepare the empty report
	outputReport := checkcompleted.New(event.Payload.CheckID)

	auth := event.Payload.AWSAuth

	notUsed := 0

	regions := awsregions.GetRegions()

	// see https://godoc.org/github.com/aws/aws-sdk-go/service/ec2#Region
	for _, region := range regions {

		svc := NewELBClient(auth, region)

		input := elb.DescribeLoadBalancersInput{}

		res, err := svc.DescribeLoadBalancers(&input)
		if err != nil {
			return nil, fmt.Errorf("could not fetch ELBs: %s", err)
		}

		for _, lb := range res.LoadBalancers {
			log.WithFields(log.Fields{
				"LoadBalancerArn": *lb.LoadBalancerArn,
			}).Debug("found load balancer")

			tgInput := elb.DescribeTargetGroupsInput{
				LoadBalancerArn: lb.LoadBalancerArn,
			}

			resTg, err := svc.DescribeTargetGroups(&tgInput)
			if err != nil {
				return nil, fmt.Errorf(
					"could not fetch tg %s details: %s",
					*lb.LoadBalancerArn,
					err,
				)
			}

			// if it have 0 target groups - it's not needed
			if len(resTg.TargetGroups) == 0 {
				log.WithFields(log.Fields{
					"LoadBalancerArn": *lb.LoadBalancerArn,
				}).Info("load balancer does not have any target groups")
				notUsed++
				continue
			}

			for _, tg := range resTg.TargetGroups {
				log.WithFields(log.Fields{
					"TargetGroupName": *tg.TargetGroupName,
					"TargetType":      *tg.TargetType,
				}).Info("found target group")

				if *tg.TargetType == "lambda" {
					log.WithFields(log.Fields{
						"TargetGroupArn": *tg.TargetGroupArn,
					}).Info("skipping Lamba target group")
					continue
				}

				healthInput := elb.DescribeTargetHealthInput{
					TargetGroupArn: tg.TargetGroupArn,
				}

				healthRes, err := svc.DescribeTargetHealth(&healthInput)
				if err != nil {
					return nil, fmt.Errorf(
						"could not fetch health details for %s: %s",
						*tg.TargetGroupArn,
						err,
					)
				}

				unhealthyCount := 0
				for _, health := range healthRes.TargetHealthDescriptions {
					if *health.TargetHealth.State != elb.TargetHealthStateEnumHealthy {
						log.WithFields(log.Fields{
							"Id":    *health.Target.Id,
							"State": *health.TargetHealth.State,
						}).Info("found unhealthy target")
						unhealthyCount++
					}
				}

				if unhealthyCount == len(healthRes.TargetHealthDescriptions) {
					log.WithFields(log.Fields{
						"TargetGroupArn": *tg.TargetGroupArn,
					}).Info("found target group without any healthy taget")
					notUsed++
					continue
				}
			}
		}

	}

	// set severity
	severity := checkcompleted.INFO
	if notUsed != 0 {
		log.WithFields(log.Fields{
			"count": notUsed,
		}).Info("found not used load balancers")
		severity = checkcompleted.ERROR
	}

	// https://aws.amazon.com/elasticloadbalancing/pricing/
	// 16.2 = ALB / month
	outputReport.Payload.Check.Impact = int(float64(notUsed) * float64(0.027*24*30))
	outputReport.Payload.Check.Severity = severity

	log.WithFields(log.Fields{
		"checkCompleted": outputReport,
	}).Info("ELB unused check finished")

	return &outputReport, nil
}
