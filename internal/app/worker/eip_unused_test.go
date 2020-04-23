package worker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (m *mockEC2Client) DescribeEmptyAddresses(*ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	// mock response/functionality

	payload := `
	{
		"Addresses": []
	}
	`

	output := &ec2.DescribeAddressesOutput{}
	err := json.Unmarshal([]byte(payload), &output)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return output, nil
}

func (m *mockEC2Client) DescribeAddresses(*ec2.DescribeAddressesInput) (*ec2.DescribeAddressesOutput, error) {
	// mock response/functionality

	payload := `
	{
		"Addresses": [
			{
				"PublicIp": "3.126.177.196",
				"Domain": "vpc",
				"AllocationId": "eipalloc-0d2232eb0a6de672c",
				"PublicIpv4Pool": "amazon"
			},
			{
				"PublicIp": "3.127.72.72",
				"Domain": "vpc",
				"AllocationId": "eipalloc-05cb653150ef928f5",
				"PublicIpv4Pool": "amazon"
			},
			{
				"PublicIp": "3.127.72.48",
				"Domain": "vpc",
				"AllocationId": "eipalloc-05cb653150ef928g5",
				"PublicIpv4Pool": "amazon",
				"AssociationId": "eipassoc-12345678",
				"InstanceId": "i-1234567890abcdef0",
				"PrivateIpAddress": "10.0.1.241"
			}
		]
	}
	`

	output := &ec2.DescribeAddressesOutput{}
	err := json.Unmarshal([]byte(payload), &output)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return output, nil
}

func TestGetAddressesCount(t *testing.T) {

	mockSvc := &mockEC2Client{}

	Addressess, _ := mockSvc.DescribeAddresses(nil)
	emptyAddressess, _ := mockSvc.DescribeEmptyAddresses(nil)

	Addresses := Addressess.Addresses
	emptyAddresses := emptyAddressess.Addresses

	tables := []struct {
		x []*ec2.Address
		y int
	}{
		{Addresses, 2},
		{emptyAddresses, 0},
	}

	for _, table := range tables {
		total := GetAddressesCount(table.x)
		if total != table.y {
			t.Errorf("Count of unused EIPs was incorrect, got: %d, want: %d.", total, table.y)
		}
	}
}
