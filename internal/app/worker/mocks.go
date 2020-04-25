package worker

import (
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

type MockEC2Client struct {
	ec2iface.EC2API
}

type CostExplorerMockClient struct {
	costexploreriface.CostExplorerAPI
}

type MockDMSClient struct {
	databasemigrationserviceiface.DatabaseMigrationServiceAPI
}
