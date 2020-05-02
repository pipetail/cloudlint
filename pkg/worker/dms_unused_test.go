package worker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
)

func (m *MockDMSClient) DescribeEmptyDMS(*databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	// mock response/functionality

	// empty response
	payload := `
	{
		"ReplicationInstances": []
	}
	`

	output := &databasemigrationservice.DescribeReplicationInstancesOutput{}
	err := json.Unmarshal([]byte(payload), &output)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return output, nil
}

func (m *MockDMSClient) DescribeDMS(*databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	// mock response/functionality

	// example payload from https://docs.aws.amazon.com/dms/latest/APIReference/API_DescribeReplicationInstances.html
	payload := `
	{
		"ReplicationInstances": [{
			"AvailabilityZone": "us-east-1c",
			"PubliclyAccessible": true,
			"ReplicationInstanceArn": "arn:aws:dms:us-east-1:123456789012:rep:PWEBBEUNOLU7VEB2OHTEH4I4GQ",
			"ReplicationInstanceClass": "dms.t2.micro",
			"ReplicationSubnetGroup": {
				"ReplicationSubnetGroupDescription": "default",
				"Subnets": [{
						"SubnetStatus": "Active",
						"SubnetIdentifier": "subnet-f6dd91af",
						"SubnetAvailabilityZone": {
							"Name": "us-east-1d"
						}
					},
					{
						"SubnetStatus": "Active",
						"SubnetIdentifier": "subnet-3605751d",
						"SubnetAvailabilityZone": {
							"Name": "us-east-1b"
						}
					},
					{
						"SubnetStatus": "Active",
						"SubnetIdentifier": "subnet-c2daefb5",
						"SubnetAvailabilityZone": {
							"Name": "us-east-1c"
						}
					},
					{
						"SubnetStatus": "Active",
						"SubnetIdentifier": "subnet-85e90cb8",
						"SubnetAvailabilityZone": {
							"Name": "us-east-1e"
						}
					}
				],
				"VpcId": "vpc-6741a603",
				"SubnetGroupStatus": "Complete",
				"ReplicationSubnetGroupIdentifier": "default"
			},
			"AutoMinorVersionUpgrade": true,
			"MultiAZ": true,
			"ReplicationInstanceStatus": "creating",
			"KmsKeyId": "arn:aws:kms:us-east-1:123456789012:key/4dc17316-5543-4ded-b1e3-d53a7cfb411d",
			"AllocatedStorage": 5,
			"EngineVersion": "1.5.0",
			"ReplicationInstanceIdentifier": "test-rep-1",
			"PreferredMaintenanceWindow": "sun:06:00-sun:14:00",
			"PendingModifiedValues": {

			}
		}]
	}
	`

	output := &databasemigrationservice.DescribeReplicationInstancesOutput{}
	err := json.Unmarshal([]byte(payload), &output)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return output, nil
}

func TestGetDmsStoragePriceInRegion(t *testing.T) {

	mockSvc := &MockDMSClient{}

	DMSs, _ := mockSvc.DescribeDMS(nil)
	emptyDMSs, _ := mockSvc.DescribeEmptyDMS(nil)

	DMS := DMSs.ReplicationInstances
	emptyDMS := emptyDMSs.ReplicationInstances

	tables := []struct {
		x []*databasemigrationservice.ReplicationInstance
		y int
	}{
		{DMS, 1},
		{emptyDMS, 0},
	}

	for _, table := range tables {

		for _, instance := range table.x {
			total := int(getDmsStoragePriceInRegion("us-east-1", instance))
			if total != table.y {
				t.Errorf("Price of DMS replication instances was incorrect, got: %d, want: %d.", total, table.y)
			}
		}
	}
}
