package awsregions

// GetRegions returns all the current checks we support
func GetRegions() []string {

	// // Create new EC2 client to get all regions
	// ec2client := ec2.New(sess, &aws.Config{Credentials: creds})

	// regions, err := ec2client.DescribeRegions(&ec2.DescribeRegionsInput{})

	// if err != nil {
	// 	panic(err)
	// }

	// // some regions are special, just like you are
	// if *region.OptInStatus == "not-opted-in" {
	// 	continue
	// }

	regions := []string{
		"eu-north-1",
		"ap-south-1",
		"eu-west-3",
		"eu-west-2",
		"eu-west-1",
		"ap-northeast-2",
		"ap-northeast-1",
		"sa-east-1",
		"ca-central-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"eu-central-1",
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
	}

	return regions
}
