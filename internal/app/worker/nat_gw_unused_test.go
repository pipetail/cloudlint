package worker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (m *mockEC2Client) DescribeEmptyNatGateways(*ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	// mock response/functionality

	// empty response
	//  aws  ec2 describe-vpc-endpoints
	// {
	//	"NatGateways": []
	// }
	NatGatewaysOutput := &ec2.DescribeNatGatewaysOutput{
		NatGateways: []*ec2.NatGateway{},
	}

	return NatGatewaysOutput, nil
	//return nil, nil
}

func (m *mockEC2Client) DescribeNatGateways(*ec2.DescribeNatGatewaysInput) (*ec2.DescribeNatGatewaysOutput, error) {
	// mock response/functionality

	// example payload from https://docs.aws.amazon.com/cli/latest/reference/ec2/describe-nat-gateways.html
	payload := `
	{
		"NatGateways": [
		  {
			"NatGatewayAddresses": [
			  {
				"PublicIp": "198.11.222.333",
				"NetworkInterfaceId": "eni-9dec76cd",
				"AllocationId": "eipalloc-89c620ec",
				"PrivateIp": "10.0.0.149"
			  }
			],
			"VpcId": "vpc-1a2b3c4d",
			"Tags": [
					{
						"Value": "IT",
						"Key": "Department"
					}
			],
			"State": "available",
			"NatGatewayId": "nat-05dba92075d71c408",
			"SubnetId": "subnet-847e4dc2",
			"CreateTime": "2015-12-01T12:26:55.983Z"
		  },
		  {
			"NatGatewayAddresses": [
			  {
				"PublicIp": "1.2.3.12",
				"NetworkInterfaceId": "eni-71ec7621",
				"AllocationId": "eipalloc-5d42583f",
				"PrivateIp": "10.0.0.77"
			  }
			],
			"VpcId": "vpc-11aa22bb",
			"Tags": [
					{
						"Value": "Finance",
						"Key": "Department"
					}
			],
			"State": "available",
			"NatGatewayId": "nat-0a93acc57881d4199",
			"SubnetId": "subnet-7f7e4d39",
			"CreateTime": "2015-12-01T12:09:22.040Z"
		  }
		]
	  }
	`

	natGatewayOutput := &ec2.DescribeNatGatewaysOutput{}
	err := json.Unmarshal([]byte(payload), &natGatewayOutput)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return natGatewayOutput, nil
	//return nil, nil
}

func TestNatGatewayUnusedCount(t *testing.T) {

	mockSvc := &mockEC2Client{}

	NatGatewayss, _ := mockSvc.DescribeNatGateways(nil)
	emptyNatGatewayss, _ := mockSvc.DescribeEmptyNatGateways(nil)

	NatGateways := NatGatewayss.NatGateways
	emptyNatGateways := emptyNatGatewayss.NatGateways

	tables := []struct {
		x []*ec2.NatGateway
		y int
	}{
		{NatGateways, 2},
		{emptyNatGateways, 0},
	}

	for _, table := range tables {
		total := GetNatGatewaysCount(table.x)
		if total != table.y {
			t.Errorf("Count of nat gateways was incorrect, got: %d, want: %d.", total, table.y)
		}
	}
}
