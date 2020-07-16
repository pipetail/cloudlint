package worker

import (
	"encoding/json"
	"fmt"
    "github.com/aws/aws-sdk-go/service/pricing"
    "math"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

func (m *MockEC2Client) DescribeVolumes2(*ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	payload := `
	{
		"Volumes": [
			{
				"AvailabilityZone": "us-east-1a",
				"Attachments": [
					{
						"AttachTime": "2013-12-18T22:35:00.000Z",
						"InstanceId": "i-1234567890abcdef0",
						"VolumeId": "vol-049df61146c4d7901",
						"State": "attached",
						"DeleteOnTermination": true,
						"Device": "/dev/sda1"
					}
				],
				"Encrypted": false,
				"VolumeType": "gp2",
				"VolumeId": "vol-049df61146c4d7901",
				"State": "in-use",
				"SnapshotId": "snap-1234567890abcdef0",
				"CreateTime": "2013-12-18T22:35:00.084Z",
				"Size": 8
			},
			{
				"AvailabilityZone": "us-east-1a",
				"Attachments": [],
				"Encrypted": false,
				"VolumeType": "gp2",
				"VolumeId": "vol-1234567890abcdef0",
				"State": "available",
				"Iops": 1000,
				"SnapshotId": null,
				"CreateTime": "2014-02-27T00:02:41.791Z",
				"Size": 100
			}
		]
	}
`
	output := &ec2.DescribeVolumesOutput{}
	err := json.Unmarshal([]byte(payload), &output)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return output, nil
}

func (m *MockEC2Client) DescribeVolumes(input *ec2.DescribeVolumesInput) (*ec2.DescribeVolumesOutput, error) {
	// mock response/functionality

	volumeID := "vol-0244f4fb5eb3e997e"
	tagKey := "env"
	tagValue := "prod"
	tag := ec2.Tag{Key: &tagKey, Value: &tagValue}
	volumeSize := int64(30)
	volumeType := "gp2"

	attachTime, _ := time.Parse("2019-12-02T17:21:17.000Z", "2019-12-02T17:21:17.000Z")
	createTime, _ := time.Parse("2019-12-02T17:21:17.000Z", "2019-12-02T17:21:17.875Z")

	volumesOutput := &ec2.DescribeVolumesOutput{
		Volumes: []*ec2.Volume{
			{
				VolumeId:   &volumeID,
				Tags:       []*ec2.Tag{&tag},
				Size:       &volumeSize,
				VolumeType: &volumeType,
			},
			{
				AvailabilityZone: aws.String("eu-central-1c"),
				Attachments: []*ec2.VolumeAttachment{
					{
						AttachTime:          aws.Time(attachTime),
						InstanceId:          aws.String("i-0c4e26ce2b59e3517"),
						VolumeId:            aws.String("vol-01dcff870b37e528e"),
						State:               aws.String("attached"),
						DeleteOnTermination: aws.Bool(true),
						Device:              aws.String("/dev/sda1"),
					},
				},
				Tags: []*ec2.Tag{
					{
						Value: aws.String("Gitlab Runner"),
						Key:   aws.String("Name"),
					},
				},
				Encrypted:  aws.Bool(false),
				VolumeType: aws.String("gp2"),
				VolumeId:   aws.String("vol-01dcff870b37e528e"),
				State:      aws.String("in-use"),
				Iops:       aws.Int64(100),
				SnapshotId: aws.String("snap-0cbf006a93852fd22"),
				CreateTime: aws.Time(createTime),
				Size:       aws.Int64(20),
			},
			{
				AvailabilityZone: aws.String("eu-central-1b"),
				Attachments: []*ec2.VolumeAttachment{
					{
						AttachTime:          aws.Time(attachTime),
						InstanceId:          aws.String("i-0c4e26ce2b48e3517"),
						VolumeId:            aws.String("vol-01dcff870437e528e"),
						State:               aws.String("attached"),
						DeleteOnTermination: aws.Bool(true),
						Device:              aws.String("/dev/sda2"),
					},
				},
				Tags: []*ec2.Tag{
					{
						Value: aws.String("EKS node"),
						Key:   aws.String("Name"),
					},
				},
				Encrypted:  aws.Bool(false),
				VolumeType: aws.String("io1"),
				VolumeId:   aws.String("vol-01dcff870437e528e"),
				State:      aws.String("in-use"),
				Iops:       aws.Int64(1600),
				SnapshotId: aws.String("snap-0cbf006a93898fd22"),
				CreateTime: aws.Time(createTime),
				Size:       aws.Int64(200),
			},
		},
	}

	if m.region == "eu-central-1" {
		return m.DescribeVolumes2(input)
	}

	return volumesOutput, nil
}

func (m *MockPricingClient) GetProducts(*pricing.GetProductsInput) (*pricing.GetProductsOutput, error) {
    priceList := make([]aws.JSONValue, 0)

    terms := `{
		"terms": {
			"OnDemand": {
				"XUZVGXG9M6A44GDS.JRTCKXETXF": {
					"effectiveDate": "2020-06-01T00:00:00Z",
					"offerTermCode": "JRTCKXETXF",
					"priceDimensions": {
						"XUZVGXG9M6A44GDS.JRTCKXETXF.6YS6EN2CT7": {
							"pricePerUnit": {
								"USD": "0.119"
							},
							"unit": "GB-Mo"
						}
					},
					"sku": "XUZVGXG9M6A44GDS",
					"termAttributes": {}
				}
			}
		},
		"version": "20200618221809"
	}`

    price := &aws.JSONValue{}
    if err := json.Unmarshal([]byte(terms), &price); err != nil {
        panic("failed to unmarshal JSONValue, " + err.Error())
    }

    priceList = append(priceList, *price)

    productsOutput := &pricing.GetProductsOutput{
        FormatVersion: aws.String("aws_v1"),
        PriceList: priceList,
    }

    return productsOutput, nil
}

var Eps float64 = 0.00000001

func FloatEquals(a, b float64) bool {
	if math.Abs(a-b) < Eps {
		return true
	}
	return false
}

func TestFilterDetachedVolumes(t *testing.T) {

	mockSvc := &MockEC2Client{region: "us-east-1"}
	mockSvc2 := &MockEC2Client{region: "eu-central-1"}

	volumess, _ := mockSvc.DescribeVolumes(nil)
	volumess2, _ := mockSvc2.DescribeVolumes(nil)

	volumes := volumess.Volumes
	volumes2 := volumess2.Volumes

	//fmt.Printf("volume: %v", volumes)

	tables := []struct {
		x []*ec2.Volume
		y int
	}{
		{volumes, 1},
		{volumes2, 1},
	}

	for _, table := range tables {
		total := len(filterDetachedVolumes(table.x))
		if total != table.y {
			t.Errorf("filterDetachedVolumes length was incorrect, got: %d, want: %d.", total, table.y)
		}
	}
}

func TestGetVolumesPrice(t *testing.T) {

    region1 := "us-east-1"
    region2 := "eu-central-1"

    mockSvc := &MockEC2Client{region: region1}
    mockSvc2 := &MockEC2Client{region: region2}

	mockPricing := &MockPricingClient{}

	volumess, _ := mockSvc.DescribeVolumes(nil)
	volumess2, _ := mockSvc2.DescribeVolumes(nil)

	volumes := filterDetachedVolumes(volumess.Volumes)
	volumes2 := filterDetachedVolumes(volumess2.Volumes)

	tables := []struct {
		x []*ec2.Volume
		y float64
		region string
	}{
		{volumes, 3.57, region1},
		{volumes2, 11.9, region2},
	}

	for _, table := range tables {
		total := GetVolumesPrice(table.x, mockPricing, table.region)
		if !FloatEquals(total, table.y) {
			t.Errorf("TotalPrice of volumes was incorrect, got: %f, want: %f.", total, table.y)
		}
	}
}

func TestGetVolumesWithinRegion(t *testing.T) {

	mockSvc := &MockEC2Client{region: "eu-central-1"}
	mockSvc2 := &MockEC2Client{region: "us-east-1"}

	tables := []struct {
		x ec2iface.EC2API
		y int
	}{
		{mockSvc, 2},
		{mockSvc2, 3},
	}

	for _, table := range tables {
		total := len(getVolumesWithinRegion(table.x))
		if total != table.y {
			t.Errorf("TotalPrice of volumes was incorrect, got: %d, want: %d.", total, table.y)
		}
	}
}
