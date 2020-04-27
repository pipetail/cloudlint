package worker

import (
	"github.com/aws/aws-sdk-go/service/costexplorer/costexploreriface"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// MockEC2Client mocks EC2 API client for unit testing purposes
type MockEC2Client struct {
	ec2iface.EC2API
}

// CostExplorerMockClient mocks CostExplorer API client for unit testing purposes
type CostExplorerMockClient struct {
	costexploreriface.CostExplorerAPI
}

// MockDMSClient mocks DMS API client for unit testing purposes
type MockDMSClient struct {
	databasemigrationserviceiface.DatabaseMigrationServiceAPI
}
