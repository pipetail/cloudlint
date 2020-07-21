package worker

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/pricing"
)

func TestExtractPrice(t *testing.T) {

	mockSvc := &MockPricingClient{}

	products, _ := mockSvc.GetProducts(nil)

	tables := []struct {
		x *pricing.GetProductsOutput
		y ec2price
	}{
		{products, ec2price{0.096, "Hrs"}},
	}

	for _, table := range tables {
		total := extractPrice(table.x)
		if total.value != table.y.value {
			t.Errorf("Price value was incorrect, got: %f, want: %f.", total.value, table.y.value)
		}
		if total.unit != table.y.unit {
			t.Errorf("Price unit was incorrect, got: %s, want: %s.", total.unit, table.y.unit)
		}
	}
}

func TestGetPricePerMonth(t *testing.T) {

	tables := []struct {
		x ec2price
		y float64
	}{
		{ec2price{0.096, "Hrs"}, 69.12},
		//{ec2price{0.96, "Hrs"}, 691.200000},
		{ec2price{0, "Hrs"}, 0},
		{ec2price{-1, "Hrs"}, 0},
	}

	for _, table := range tables {
		total, err := getPricePerMonth(table.x)
		if err != nil {
			t.Errorf("Got error when calling getPricePerMonth")
		}
		if total != table.y {
			t.Errorf("Price per month value was incorrect, got: %f, want: %f.", total, table.y)
		}
	}
}

func TestGetPricePerMonthError(t *testing.T) {

	tables := []struct {
		x ec2price
		y float64
	}{
		{ec2price{0.96, "USD"}, 69.12},
	}

	for _, table := range tables {
		_, err := getPricePerMonth(table.x)
		if err == nil {
			t.Errorf("GetPricePerMonth should error due to wrong unit but it didn't.")
		}
	}
}

func (m *MockPricingClient) GetProducts(*pricing.GetProductsInput) (*pricing.GetProductsOutput, error) {
	// mock response/functionality

	payload := `
	{
		"FormatVersion": "aws_v1",
		"NextToken": null,
		"PriceList": [
			{
				"product": {
					"attributes": {
						"capacitystatus": "Used",
						"clockSpeed": "3.1 GHz",
						"currentGeneration": "Yes",
						"dedicatedEbsThroughput": "Up to 2120 Mbps",
						"ecu": "10",
						"enhancedNetworkingSupported": "Yes",
						"instanceFamily": "General purpose",
						"instanceType": "m5.large",
						"intelAvx2Available": "Yes",
						"intelAvxAvailable": "Yes",
						"intelTurboAvailable": "Yes",
						"licenseModel": "No License required",
						"location": "US East (N. Virginia)",
						"locationType": "AWS Region",
						"memory": "8 GiB",
						"networkPerformance": "Up to 10 Gigabit",
						"normalizationSizeFactor": "4",
						"operatingSystem": "Linux",
						"operation": "RunInstances",
						"physicalProcessor": "Intel Xeon Platinum 8175 (Skylake)",
						"preInstalledSw": "NA",
						"processorArchitecture": "64-bit",
						"processorFeatures": "Intel AVX; Intel AVX2; Intel AVX512; Intel Turbo",
						"servicecode": "AmazonEC2",
						"servicename": "Amazon Elastic Compute Cloud",
						"storage": "EBS only",
						"tenancy": "Shared",
						"usagetype": "BoxUsage:m5.large",
						"vcpu": "2"
					},
					"productFamily": "Compute Instance",
					"sku": "6C86BEPQVG73ZGGR"
				},
				"publicationDate": "2020-07-09T22:30:13Z",
				"serviceCode": "AmazonEC2",
				"terms": {
					"OnDemand": {
						"6C86BEPQVG73ZGGR.JRTCKXETXF": {
							"effectiveDate": "2020-07-01T00:00:00Z",
							"offerTermCode": "JRTCKXETXF",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.JRTCKXETXF.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "$0.096 per On Demand Linux m5.large Instance Hour",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0960000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.JRTCKXETXF.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {}
						}
					},
					"Reserved": {
						"6C86BEPQVG73ZGGR.38NPMPTW36": {
							"effectiveDate": "2020-04-01T00:00:00Z",
							"offerTermCode": "38NPMPTW36",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.38NPMPTW36.2TG2D8R56U": {
									"appliesTo": [],
									"description": "Upfront Fee",
									"pricePerUnit": {
										"USD": "505"
									},
									"rateCode": "6C86BEPQVG73ZGGR.38NPMPTW36.2TG2D8R56U",
									"unit": "Quantity"
								},
								"6C86BEPQVG73ZGGR.38NPMPTW36.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0190000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.38NPMPTW36.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "3yr",
								"OfferingClass": "standard",
								"PurchaseOption": "Partial Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.4NA7Y494T4": {
							"effectiveDate": "2020-04-01T00:00:00Z",
							"offerTermCode": "4NA7Y494T4",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.4NA7Y494T4.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0600000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.4NA7Y494T4.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "1yr",
								"OfferingClass": "standard",
								"PurchaseOption": "No Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.6QCMYABX3D": {
							"effectiveDate": "2020-04-01T00:00:00Z",
							"offerTermCode": "6QCMYABX3D",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.6QCMYABX3D.2TG2D8R56U": {
									"appliesTo": [],
									"description": "Upfront Fee",
									"pricePerUnit": {
										"USD": "494"
									},
									"rateCode": "6C86BEPQVG73ZGGR.6QCMYABX3D.2TG2D8R56U",
									"unit": "Quantity"
								},
								"6C86BEPQVG73ZGGR.6QCMYABX3D.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "USD 0.0 per Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0000000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.6QCMYABX3D.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "1yr",
								"OfferingClass": "standard",
								"PurchaseOption": "All Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.7NE97W5U4E": {
							"effectiveDate": "2017-10-31T23:59:59Z",
							"offerTermCode": "7NE97W5U4E",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.7NE97W5U4E.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0710000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.7NE97W5U4E.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "1yr",
								"OfferingClass": "convertible",
								"PurchaseOption": "No Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.BPH4J8HBKS": {
							"effectiveDate": "2020-04-01T00:00:00Z",
							"offerTermCode": "BPH4J8HBKS",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.BPH4J8HBKS.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0410000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.BPH4J8HBKS.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "3yr",
								"OfferingClass": "standard",
								"PurchaseOption": "No Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.CUZHX8X6JH": {
							"effectiveDate": "2017-10-31T23:59:59Z",
							"offerTermCode": "CUZHX8X6JH",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.CUZHX8X6JH.2TG2D8R56U": {
									"appliesTo": [],
									"description": "Upfront Fee",
									"pricePerUnit": {
										"USD": "294"
									},
									"rateCode": "6C86BEPQVG73ZGGR.CUZHX8X6JH.2TG2D8R56U",
									"unit": "Quantity"
								},
								"6C86BEPQVG73ZGGR.CUZHX8X6JH.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0340000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.CUZHX8X6JH.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "1yr",
								"OfferingClass": "convertible",
								"PurchaseOption": "Partial Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.HU7G6KETJZ": {
							"effectiveDate": "2020-04-01T00:00:00Z",
							"offerTermCode": "HU7G6KETJZ",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.HU7G6KETJZ.2TG2D8R56U": {
									"appliesTo": [],
									"description": "Upfront Fee",
									"pricePerUnit": {
										"USD": "252"
									},
									"rateCode": "6C86BEPQVG73ZGGR.HU7G6KETJZ.2TG2D8R56U",
									"unit": "Quantity"
								},
								"6C86BEPQVG73ZGGR.HU7G6KETJZ.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0290000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.HU7G6KETJZ.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "1yr",
								"OfferingClass": "standard",
								"PurchaseOption": "Partial Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.MZU6U2429S": {
							"effectiveDate": "2017-10-31T23:59:59Z",
							"offerTermCode": "MZU6U2429S",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.MZU6U2429S.2TG2D8R56U": {
									"appliesTo": [],
									"description": "Upfront Fee",
									"pricePerUnit": {
										"USD": "1161"
									},
									"rateCode": "6C86BEPQVG73ZGGR.MZU6U2429S.2TG2D8R56U",
									"unit": "Quantity"
								},
								"6C86BEPQVG73ZGGR.MZU6U2429S.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0000000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.MZU6U2429S.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "3yr",
								"OfferingClass": "convertible",
								"PurchaseOption": "All Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.NQ3QZPMQV9": {
							"effectiveDate": "2020-04-01T00:00:00Z",
							"offerTermCode": "NQ3QZPMQV9",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.NQ3QZPMQV9.2TG2D8R56U": {
									"appliesTo": [],
									"description": "Upfront Fee",
									"pricePerUnit": {
										"USD": "949"
									},
									"rateCode": "6C86BEPQVG73ZGGR.NQ3QZPMQV9.2TG2D8R56U",
									"unit": "Quantity"
								},
								"6C86BEPQVG73ZGGR.NQ3QZPMQV9.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "USD 0.0 per Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0000000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.NQ3QZPMQV9.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "3yr",
								"OfferingClass": "standard",
								"PurchaseOption": "All Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.R5XV2EPZQZ": {
							"effectiveDate": "2017-10-31T23:59:59Z",
							"offerTermCode": "R5XV2EPZQZ",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.R5XV2EPZQZ.2TG2D8R56U": {
									"appliesTo": [],
									"description": "Upfront Fee",
									"pricePerUnit": {
										"USD": "592"
									},
									"rateCode": "6C86BEPQVG73ZGGR.R5XV2EPZQZ.2TG2D8R56U",
									"unit": "Quantity"
								},
								"6C86BEPQVG73ZGGR.R5XV2EPZQZ.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0230000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.R5XV2EPZQZ.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "3yr",
								"OfferingClass": "convertible",
								"PurchaseOption": "Partial Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.VJWZNREJX2": {
							"effectiveDate": "2017-10-31T23:59:59Z",
							"offerTermCode": "VJWZNREJX2",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.VJWZNREJX2.2TG2D8R56U": {
									"appliesTo": [],
									"description": "Upfront Fee",
									"pricePerUnit": {
										"USD": "577"
									},
									"rateCode": "6C86BEPQVG73ZGGR.VJWZNREJX2.2TG2D8R56U",
									"unit": "Quantity"
								},
								"6C86BEPQVG73ZGGR.VJWZNREJX2.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0000000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.VJWZNREJX2.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "1yr",
								"OfferingClass": "convertible",
								"PurchaseOption": "All Upfront"
							}
						},
						"6C86BEPQVG73ZGGR.Z2E3P23VKM": {
							"effectiveDate": "2017-10-31T23:59:59Z",
							"offerTermCode": "Z2E3P23VKM",
							"priceDimensions": {
								"6C86BEPQVG73ZGGR.Z2E3P23VKM.6YS6EN2CT7": {
									"appliesTo": [],
									"beginRange": "0",
									"description": "Linux/UNIX (Amazon VPC), m5.large reserved instance applied",
									"endRange": "Inf",
									"pricePerUnit": {
										"USD": "0.0490000000"
									},
									"rateCode": "6C86BEPQVG73ZGGR.Z2E3P23VKM.6YS6EN2CT7",
									"unit": "Hrs"
								}
							},
							"sku": "6C86BEPQVG73ZGGR",
							"termAttributes": {
								"LeaseContractLength": "3yr",
								"OfferingClass": "convertible",
								"PurchaseOption": "No Upfront"
							}
						}
					}
				},
				"version": "20200709223013"
			}
		]
	}
	`

	output := &pricing.GetProductsOutput{}
	err := json.Unmarshal([]byte(payload), &output)
	if err != nil {
		fmt.Printf("could not parse input json: %s", err)
	}

	return output, nil
}
