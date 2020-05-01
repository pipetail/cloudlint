package worker

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-sdk-go/service/ec2"
)

func (m *MockEC2Client) DescribeImages(*ec2.DescribeImagesInput) (*ec2.DescribeImagesOutput, error) {
	payload := `{
		"Images": [
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-02T13:36:12.000Z",
				"ImageId": "ami-0590f6463f77936cd",
				"ImageLocation": "101632338718/packer-ds-1583155951",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-085f64be6b0ddbc34",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583155951",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-11T13:11:24.000Z",
				"ImageId": "ami-05c13b355a7df6e6e",
				"ImageLocation": "101632338718/packer-ds-1583931973",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-03f8ab627895c4da3",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583931973",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-04T09:55:37.000Z",
				"ImageId": "ami-085057fc8b6b02831",
				"ImageLocation": "101632338718/packer-ds-1583315486",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-08230a89a1656bedd",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583315486",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-11T12:01:55.000Z",
				"ImageId": "ami-085e96e019a3733cc",
				"ImageLocation": "101632338718/packer-ds-1583927867",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-0649919c19e4c51f2",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583927867",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-09T08:13:24.000Z",
				"ImageId": "ami-09c1aaa3dc19b7f39",
				"ImageLocation": "101632338718/packer-ds-1583741378",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-09740ead1923f2bd7",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583741378",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-09T08:53:55.000Z",
				"ImageId": "ami-0a7cc784f2559208a",
				"ImageLocation": "101632338718/packer-ds-1583743820",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-06ed14f768aa12384",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583743820",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-04T10:07:18.000Z",
				"ImageId": "ami-0b524c123edd64f4c",
				"ImageLocation": "101632338718/packer-ds-1583316166",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-07c5035427dc07653",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583316166",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-03T07:33:55.000Z",
				"ImageId": "ami-0c631bf4eb699e2bc",
				"ImageLocation": "101632338718/packer-ds-1583220579",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-0e47a9f2dc70d1dcf",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583220579",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-04T11:33:33.000Z",
				"ImageId": "ami-0d005d8c2e130596e",
				"ImageLocation": "101632338718/packer-ds-1583321355",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-00ba11550d69f5d9a",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583321355",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-04T11:56:17.000Z",
				"ImageId": "ami-0d36fdb34fa667ee8",
				"ImageLocation": "101632338718/packer-ds-1583322737",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-0f08946b15663d179",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583322737",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-02T14:31:47.000Z",
				"ImageId": "ami-0e3aeec8d22e2bbf2",
				"ImageLocation": "101632338718/packer-ds-1583159298",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-08d2196789112a8a8",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583159298",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-03T10:04:09.000Z",
				"ImageId": "ami-0f07abc4fc32ffb06",
				"ImageLocation": "101632338718/packer-ds-1583229592",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-0aa8a38c65d4bb30e",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1583229592",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			},
			{
				"Architecture": "x86_64",
				"CreationDate": "2020-03-16T12:30:09.000Z",
				"ImageId": "ami-0fbe38ae62f27a845",
				"ImageLocation": "101632338718/packer-ds-1584361535",
				"ImageType": "machine",
				"Public": false,
				"OwnerId": "101632338718",
				"State": "available",
				"BlockDeviceMappings": [
					{
						"DeviceName": "/dev/sda1",
						"Ebs": {
							"DeleteOnTermination": true,
							"SnapshotId": "snap-0e7280f4c08584980",
							"VolumeSize": 8,
							"VolumeType": "gp2",
							"Encrypted": false
						}
					},
					{
						"DeviceName": "/dev/sdb",
						"VirtualName": "ephemeral0"
					},
					{
						"DeviceName": "/dev/sdc",
						"VirtualName": "ephemeral1"
					}
				],
				"EnaSupport": true,
				"Hypervisor": "xen",
				"Name": "packer-ds-1584361535",
				"RootDeviceName": "/dev/sda1",
				"RootDeviceType": "ebs",
				"SriovNetSupport": "simple",
				"VirtualizationType": "hvm"
			}
		]
	}`

	imagesOutput := &ec2.DescribeImagesOutput{}
	err := json.Unmarshal([]byte(payload), &imagesOutput)
	if err != nil {
		return nil, err
	}

	return imagesOutput, nil
}

func TestAmisUnusedCount(t *testing.T) {
	mockSvc := &MockEC2Client{}
	_, _ = mockSvc.DescribeImages(nil)
	//tbd
}
