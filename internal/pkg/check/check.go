package check

import (
	uuid "github.com/google/uuid"
	"github.com/pipetail/cloudlint/internal/pkg/checkreport"
)

/*
{
    "name": "Check",
    "payload": {
        "awsAuth": {
            "roleArn": "<String>",
            "externalId": "<String>"
        },
        "reportId": "<Uuid>",
        "checkId": "<Uuid>",
        "checkType": "<String>"
    }
}

*/

// Check struct with all the information we will show on the frontend
type Check struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Group       string `json:"group"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// GetChecks returns all the current checks we support
func GetChecks() []Check {

	checks := []Check{
		{
			ID:          uuid.New().String(),
			Name:        "AWS Monthly Bill",
			Group:       "General info",
			Type:        "aws_monthly_bill",
			Description: "Unblended cost for all services in cost explorer",
		},
		{
			ID:          uuid.New().String(),
			Name:        "AWS DMS Replication instances",
			Group:       "Resources with no usage",
			Type:        "aws_dms_unused",
			Description: "Total number of dms replication instances",
		},
		{
			ID:          uuid.New().String(),
			Name:        "EBS Unused",
			Group:       "Resources with no usage",
			Type:        "total_ebs_unused",
			Description: "Total unused EBS storage in GiB",
		},
		{
			ID:          uuid.New().String(),
			Name:        "AWS Paid Support",
			Group:       "Resources with no usage",
			Type:        "aws_paid_support",
			Description: "This can be bought operatively for single questions, not needed to be prepaid",
		},
		{
			ID:          uuid.New().String(),
			Name:        "EBS snapshots are old",
			Group:       "Resources with no usage",
			Type:        "aws_ebs_snapshots_old",
			Description: "Some EBS snapshots detected are very old",
		},
		{
			ID:          uuid.New().String(),
			Name:        "ELB Unused",
			Group:       "Resources with no usage",
			Type:        "total_elb_unused",
			Description: "Total unused ELBs",
		},
		{
			ID:          uuid.New().String(),
			Name:        "VPC S3 endpoints are not used",
			Group:       "Incorrect service usage",
			Type:        "vpc_eps_notused",
			Description: "VPC S3 Endpoints are not used",
		},
		{
			ID:          uuid.New().String(),
			Name:        "NAT Gateways are unused",
			Group:       "Incorrect service usage",
			Type:        "nat_gw_unused",
			Description: "NAT Gateways are alone in a subnet",
		},
		{
			ID:          uuid.New().String(),
			Name:        "Elastic IPs are unused",
			Group:       "Incorrect service usage",
			Type:        "eip_unused",
			Description: "Elastic IP addresses are unassociated to any instance",
		},
		{
			ID:          uuid.New().String(),
			Name:        "AMIs are too old",
			Group:       "Resources with no usage",
			Type:        "ami_old",
			Description: "There are possibly unused AMIs that increase S3 storage price",
		},
		{
			ID:          uuid.New().String(),
			Name:        "EC2 instances are not EBS Optimized",
			Group:       "Incorrect service usage",
			Type:        "ebs_opt",
			Description: "There are some EC2 instances with EBS attatached but at the same time they have EBS Optimization disabled",
		},
	}

	return checks
}

// Event struct for Check command
type Event struct {
	Name    string  `json:"name"`
	Payload Payload `json:"payload"`
}

// Payload struct for Check command
type Payload struct {
	AWSAuth   checkreport.AwsAuth `json:"awsAuth"`
	CheckType string              `json:"checkType"`
	CheckID   string              `json:"checkId"`
	ReportID  string              `json:"reportId"`
}

// New - constructor for creating CheckReport event
func New(reportID string) Event {
	return Event{
		Name: "Check",
		Payload: Payload{
			ReportID: reportID,
		},
	}
}
