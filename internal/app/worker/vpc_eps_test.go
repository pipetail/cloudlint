package worker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (m *mockEC2Client) DescribeEmptyVpcEndpoints(*ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	// mock response/functionality

	// empty response
	//  aws  ec2 describe-vpc-endpoints
	// {
	//	"VpcEndpoints": []
	// }
	vpcEndpointsOutput := &ec2.DescribeVpcEndpointsOutput{
		VpcEndpoints: []*ec2.VpcEndpoint{},
	}

	return vpcEndpointsOutput, nil
	//return nil, nil
}

func (m *mockEC2Client) DescribeVpcEndpointsDocs(*ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {

	payload := `
{
    "VpcEndpoints": [
        {
            "PolicyDocument": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":\"*\",\"Action\":\"*\",\"Resource\":\"*\"}]}",
            "VpcId": "vpc-aabb1122",
            "NetworkInterfaceIds": [],
            "SubnetIds": [],
            "PrivateDnsEnabled": true,
            "State": "available",
            "ServiceName": "com.amazonaws.us-east-1.dynamodb",
            "RouteTableIds": [
                "rtb-3d560345"
            ],
            "Groups": [],
            "VpcEndpointId": "vpce-032a826a",
            "VpcEndpointType": "Gateway",
            "CreationTimestamp": "2017-09-05T20:41:28Z",
            "DnsEntries": [],
            "OwnerId": "123456789012"
        },
        {
            "PolicyDocument": "{\n  \"Statement\": [\n    {\n      \"Action\": \"*\", \n      \"Effect\": \"Allow\", \n      \"Principal\": \"*\", \n      \"Resource\": \"*\"\n    }\n  ]\n}",
            "VpcId": "vpc-1a2b3c4d",
            "NetworkInterfaceIds": [
                "eni-2ec2b084",
                "eni-1b4a65cf"
            ],
            "SubnetIds": [
                "subnet-d6fcaa8d",
                "subnet-7b16de0c"
            ],
            "PrivateDnsEnabled": false,
            "State": "available",
            "ServiceName": "com.amazonaws.us-east-1.elasticloadbalancing",
            "RouteTableIds": [],
            "Groups": [
                {
                    "GroupName": "default",
                    "GroupId": "sg-54e8bf31"
                }
            ],
            "VpcEndpointId": "vpce-0f89a33420c1931d7",
            "VpcEndpointType": "Interface",
            "CreationTimestamp": "2017-09-05T17:55:27.583Z",
            "DnsEntries": [
                {
                    "HostedZoneId": "Z7HUB22UULQXV",
                    "DnsName": "vpce-0f89a33420c1931d7-bluzidnv.elasticloadbalancing.us-east-1.vpce.amazonaws.com"
                },
                {
                    "HostedZoneId": "Z7HUB22UULQXV",
                    "DnsName": "vpce-0f89a33420c1931d7-bluzidnv-us-east-1b.elasticloadbalancing.us-east-1.vpce.amazonaws.com"
                },
                {
                    "HostedZoneId": "Z7HUB22UULQXV",
                    "DnsName": "vpce-0f89a33420c1931d7-bluzidnv-us-east-1a.elasticloadbalancing.us-east-1.vpce.amazonaws.com"
                }
            ],
            "OwnerId": "123456789012"
        }
    ]
}`

	vpcEndpointsOutput := &ec2.DescribeVpcEndpointsOutput{}
	err := json.Unmarshal([]byte(payload), &vpcEndpointsOutput)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return vpcEndpointsOutput, nil
}

func (m *mockEC2Client) DescribeVpcEndpoints(*ec2.DescribeVpcEndpointsInput) (*ec2.DescribeVpcEndpointsOutput, error) {
	// mock response/functionality

	payload := `
	{
		"VpcEndpoints": [
			{
				"PolicyDocument": "{\"Version\":\"2008-10-17\",\"Statement\":[{\"Effect\":\"Allow\",\"Principal\":\"*\",\"Action\":\"*\",\"Resource\":\"*\"}]}",
				"VpcId": "vpc-50c2023a",
				"Tags": [],
				"NetworkInterfaceIds": [],
				"SubnetIds": [],
				"RequesterManaged": false,
				"PrivateDnsEnabled": false,
				"State": "available",
				"ServiceName": "com.amazonaws.eu-central-1.s3",
				"RouteTableIds": [],
				"Groups": [],
				"OwnerId": "680177765279",
				"VpcEndpointId": "vpce-0667dcb6b9a5e574f",
				"VpcEndpointType": "Gateway",
				"CreationTimestamp": "2020-03-10T14:42:10.000Z",
				"DnsEntries": []
			}
		]
	}
	`

	vpcEndpointsOutput := &ec2.DescribeVpcEndpointsOutput{}
	err := json.Unmarshal([]byte(payload), &vpcEndpointsOutput)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return vpcEndpointsOutput, nil
	//return nil, nil
}

func TestGetS3VpcEndpointsCount(t *testing.T) {

	mockSvc := &mockEC2Client{}

	vpcEndpointss, _ := mockSvc.DescribeVpcEndpoints(nil)
	vpcEndpointssDocs, _ := mockSvc.DescribeVpcEndpointsDocs(nil)
	emptyVpcEndpointss, _ := mockSvc.DescribeEmptyVpcEndpoints(nil)

	vpcEndpoints := vpcEndpointss.VpcEndpoints
	vpcEndpointsDocs := vpcEndpointssDocs.VpcEndpoints
	emptyVpcEndpoints := emptyVpcEndpointss.VpcEndpoints

	tables := []struct {
		x []*ec2.VpcEndpoint
		y int
	}{
		{vpcEndpoints, 1},
		{emptyVpcEndpoints, 0},
		{vpcEndpointsDocs, 0},
	}

	for _, table := range tables {
		total := GetS3VpcEndpointsCount(table.x)
		if total != table.y {
			t.Errorf("Count of vpc endpoints was incorrect, got: %d, want: %d.", total, table.y)
		}
	}
}
