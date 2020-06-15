package awspricing

var (
	output = []byte(`
[
	{
		"product": {
			"attributes": {
				"capacitystatus": "Used",
				"clockSpeed": "3 GHz",
				"currentGeneration": "Yes",
				"dedicatedEbsThroughput": "Up to 2250 Mbps",
				"ecu": "20",
				"enhancedNetworkingSupported": "Yes",
				"instanceFamily": "Compute optimized",
				"instanceType": "c5.xlarge",
				"intelAvx2Available": "Yes",
				"intelAvxAvailable": "Yes",
				"intelTurboAvailable": "Yes",
				"licenseModel": "No License required",
				"location": "US East (N. Virginia)",
				"locationType": "AWS Region",
				"memory": "8 GiB",
				"networkPerformance": "Up to 10 Gigabit",
				"normalizationSizeFactor": "8",
				"operatingSystem": "Linux",
				"operation": "RunInstances",
				"physicalProcessor": "Intel Xeon Platinum 8124M",
				"preInstalledSw": "NA",
				"processorArchitecture": "64-bit",
				"processorFeatures": "Intel AVX; Intel AVX2; Intel AVX512; Intel Turbo",
				"servicecode": "AmazonEC2",
				"servicename": "Amazon Elastic Compute Cloud",
				"storage": "EBS only",
				"tenancy": "Shared",
				"usagetype": "BoxUsage:c5.xlarge",
				"vcpu": "4"
			},
			"productFamily": "Compute Instance",
			"sku": "35AEEWH98DECPC35"
		},
		"publicationDate": "2020-06-18T22:18:09Z",
		"serviceCode": "AmazonEC2",
		"terms": {
			"OnDemand": {
				"35AEEWH98DECPC35.JRTCKXETXF": {
					"effectiveDate": "2020-06-01T00:00:00Z",
					"offerTermCode": "JRTCKXETXF",
					"priceDimensions": {
						"35AEEWH98DECPC35.JRTCKXETXF.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "$0.17 per On Demand Linux c5.xlarge Instance Hour",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.1700000000"
							},
							"rateCode": "35AEEWH98DECPC35.JRTCKXETXF.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {}
				}
			},
			"Reserved": {
				"35AEEWH98DECPC35.38NPMPTW36": {
					"effectiveDate": "2020-04-01T00:00:00Z",
					"offerTermCode": "38NPMPTW36",
					"priceDimensions": {
						"35AEEWH98DECPC35.38NPMPTW36.2TG2D8R56U": {
							"appliesTo": [],
							"description": "Upfront Fee",
							"pricePerUnit": {
								"USD": "867"
							},
							"rateCode": "35AEEWH98DECPC35.38NPMPTW36.2TG2D8R56U",
							"unit": "Quantity"
						},
						"35AEEWH98DECPC35.38NPMPTW36.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0330000000"
							},
							"rateCode": "35AEEWH98DECPC35.38NPMPTW36.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "3yr",
						"OfferingClass": "standard",
						"PurchaseOption": "Partial Upfront"
					}
				},
				"35AEEWH98DECPC35.4NA7Y494T4": {
					"effectiveDate": "2020-04-01T00:00:00Z",
					"offerTermCode": "4NA7Y494T4",
					"priceDimensions": {
						"35AEEWH98DECPC35.4NA7Y494T4.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.1070000000"
							},
							"rateCode": "35AEEWH98DECPC35.4NA7Y494T4.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "1yr",
						"OfferingClass": "standard",
						"PurchaseOption": "No Upfront"
					}
				},
				"35AEEWH98DECPC35.6QCMYABX3D": {
					"effectiveDate": "2020-04-01T00:00:00Z",
					"offerTermCode": "6QCMYABX3D",
					"priceDimensions": {
						"35AEEWH98DECPC35.6QCMYABX3D.2TG2D8R56U": {
							"appliesTo": [],
							"description": "Upfront Fee",
							"pricePerUnit": {
								"USD": "876"
							},
							"rateCode": "35AEEWH98DECPC35.6QCMYABX3D.2TG2D8R56U",
							"unit": "Quantity"
						},
						"35AEEWH98DECPC35.6QCMYABX3D.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "USD 0.0 per Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0000000000"
							},
							"rateCode": "35AEEWH98DECPC35.6QCMYABX3D.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "1yr",
						"OfferingClass": "standard",
						"PurchaseOption": "All Upfront"
					}
				},
				"35AEEWH98DECPC35.7NE97W5U4E": {
					"effectiveDate": "2017-10-31T23:59:59Z",
					"offerTermCode": "7NE97W5U4E",
					"priceDimensions": {
						"35AEEWH98DECPC35.7NE97W5U4E.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.1230000000"
							},
							"rateCode": "35AEEWH98DECPC35.7NE97W5U4E.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "1yr",
						"OfferingClass": "convertible",
						"PurchaseOption": "No Upfront"
					}
				},
				"35AEEWH98DECPC35.BPH4J8HBKS": {
					"effectiveDate": "2020-04-01T00:00:00Z",
					"offerTermCode": "BPH4J8HBKS",
					"priceDimensions": {
						"35AEEWH98DECPC35.BPH4J8HBKS.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0710000000"
							},
							"rateCode": "35AEEWH98DECPC35.BPH4J8HBKS.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "3yr",
						"OfferingClass": "standard",
						"PurchaseOption": "No Upfront"
					}
				},
				"35AEEWH98DECPC35.CUZHX8X6JH": {
					"effectiveDate": "2017-10-31T23:59:59Z",
					"offerTermCode": "CUZHX8X6JH",
					"priceDimensions": {
						"35AEEWH98DECPC35.CUZHX8X6JH.2TG2D8R56U": {
							"appliesTo": [],
							"description": "Upfront Fee",
							"pricePerUnit": {
								"USD": "514"
							},
							"rateCode": "35AEEWH98DECPC35.CUZHX8X6JH.2TG2D8R56U",
							"unit": "Quantity"
						},
						"35AEEWH98DECPC35.CUZHX8X6JH.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0590000000"
							},
							"rateCode": "35AEEWH98DECPC35.CUZHX8X6JH.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "1yr",
						"OfferingClass": "convertible",
						"PurchaseOption": "Partial Upfront"
					}
				},
				"35AEEWH98DECPC35.HU7G6KETJZ": {
					"effectiveDate": "2020-04-01T00:00:00Z",
					"offerTermCode": "HU7G6KETJZ",
					"priceDimensions": {
						"35AEEWH98DECPC35.HU7G6KETJZ.2TG2D8R56U": {
							"appliesTo": [],
							"description": "Upfront Fee",
							"pricePerUnit": {
								"USD": "447"
							},
							"rateCode": "35AEEWH98DECPC35.HU7G6KETJZ.2TG2D8R56U",
							"unit": "Quantity"
						},
						"35AEEWH98DECPC35.HU7G6KETJZ.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0510000000"
							},
							"rateCode": "35AEEWH98DECPC35.HU7G6KETJZ.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "1yr",
						"OfferingClass": "standard",
						"PurchaseOption": "Partial Upfront"
					}
				},
				"35AEEWH98DECPC35.MZU6U2429S": {
					"effectiveDate": "2017-10-31T23:59:59Z",
					"offerTermCode": "MZU6U2429S",
					"priceDimensions": {
						"35AEEWH98DECPC35.MZU6U2429S.2TG2D8R56U": {
							"appliesTo": [],
							"description": "Upfront Fee",
							"pricePerUnit": {
								"USD": "1957"
							},
							"rateCode": "35AEEWH98DECPC35.MZU6U2429S.2TG2D8R56U",
							"unit": "Quantity"
						},
						"35AEEWH98DECPC35.MZU6U2429S.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0000000000"
							},
							"rateCode": "35AEEWH98DECPC35.MZU6U2429S.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "3yr",
						"OfferingClass": "convertible",
						"PurchaseOption": "All Upfront"
					}
				},
				"35AEEWH98DECPC35.NQ3QZPMQV9": {
					"effectiveDate": "2020-04-01T00:00:00Z",
					"offerTermCode": "NQ3QZPMQV9",
					"priceDimensions": {
						"35AEEWH98DECPC35.NQ3QZPMQV9.2TG2D8R56U": {
							"appliesTo": [],
							"description": "Upfront Fee",
							"pricePerUnit": {
								"USD": "1629"
							},
							"rateCode": "35AEEWH98DECPC35.NQ3QZPMQV9.2TG2D8R56U",
							"unit": "Quantity"
						},
						"35AEEWH98DECPC35.NQ3QZPMQV9.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "USD 0.0 per Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0000000000"
							},
							"rateCode": "35AEEWH98DECPC35.NQ3QZPMQV9.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "3yr",
						"OfferingClass": "standard",
						"PurchaseOption": "All Upfront"
					}
				},
				"35AEEWH98DECPC35.R5XV2EPZQZ": {
					"effectiveDate": "2017-10-31T23:59:59Z",
					"offerTermCode": "R5XV2EPZQZ",
					"priceDimensions": {
						"35AEEWH98DECPC35.R5XV2EPZQZ.2TG2D8R56U": {
							"appliesTo": [],
							"description": "Upfront Fee",
							"pricePerUnit": {
								"USD": "999"
							},
							"rateCode": "35AEEWH98DECPC35.R5XV2EPZQZ.2TG2D8R56U",
							"unit": "Quantity"
						},
						"35AEEWH98DECPC35.R5XV2EPZQZ.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0380000000"
							},
							"rateCode": "35AEEWH98DECPC35.R5XV2EPZQZ.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "3yr",
						"OfferingClass": "convertible",
						"PurchaseOption": "Partial Upfront"
					}
				},
				"35AEEWH98DECPC35.VJWZNREJX2": {
					"effectiveDate": "2017-10-31T23:59:59Z",
					"offerTermCode": "VJWZNREJX2",
					"priceDimensions": {
						"35AEEWH98DECPC35.VJWZNREJX2.2TG2D8R56U": {
							"appliesTo": [],
							"description": "Upfront Fee",
							"pricePerUnit": {
								"USD": "1007"
							},
							"rateCode": "35AEEWH98DECPC35.VJWZNREJX2.2TG2D8R56U",
							"unit": "Quantity"
						},
						"35AEEWH98DECPC35.VJWZNREJX2.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0000000000"
							},
							"rateCode": "35AEEWH98DECPC35.VJWZNREJX2.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "1yr",
						"OfferingClass": "convertible",
						"PurchaseOption": "All Upfront"
					}
				},
				"35AEEWH98DECPC35.Z2E3P23VKM": {
					"effectiveDate": "2017-10-31T23:59:59Z",
					"offerTermCode": "Z2E3P23VKM",
					"priceDimensions": {
						"35AEEWH98DECPC35.Z2E3P23VKM.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "Linux/UNIX (Amazon VPC), c5.xlarge reserved instance applied",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0820000000"
							},
							"rateCode": "35AEEWH98DECPC35.Z2E3P23VKM.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "35AEEWH98DECPC35",
					"termAttributes": {
						"LeaseContractLength": "3yr",
						"OfferingClass": "convertible",
						"PurchaseOption": "No Upfront"
					}
				}
			}
		},
		"version": "20200618221809"
	},
	{
		"product": {
			"attributes": {
				"capacitystatus": "UnusedCapacityReservation",
				"clockSpeed": "3 GHz",
				"currentGeneration": "Yes",
				"dedicatedEbsThroughput": "Up to 2250 Mbps",
				"ecu": "20",
				"enhancedNetworkingSupported": "Yes",
				"instanceFamily": "Compute optimized",
				"instanceType": "c5.xlarge",
				"instancesku": "35AEEWH98DECPC35",
				"intelAvx2Available": "Yes",
				"intelAvxAvailable": "Yes",
				"intelTurboAvailable": "Yes",
				"licenseModel": "No License required",
				"location": "US East (N. Virginia)",
				"locationType": "AWS Region",
				"memory": "8 GiB",
				"networkPerformance": "Up to 10 Gigabit",
				"normalizationSizeFactor": "8",
				"operatingSystem": "Linux",
				"operation": "RunInstances",
				"physicalProcessor": "Intel Xeon Platinum 8124M",
				"preInstalledSw": "NA",
				"processorArchitecture": "64-bit",
				"processorFeatures": "Intel AVX; Intel AVX2; Intel AVX512; Intel Turbo",
				"servicecode": "AmazonEC2",
				"servicename": "Amazon Elastic Compute Cloud",
				"storage": "EBS only",
				"tenancy": "Shared",
				"usagetype": "UnusedBox:c5.xlarge",
				"vcpu": "4"
			},
			"productFamily": "Compute Instance",
			"sku": "XUZVGXG9M6A44GDS"
		},
		"publicationDate": "2020-06-18T22:18:09Z",
		"serviceCode": "AmazonEC2",
		"terms": {
			"OnDemand": {
				"XUZVGXG9M6A44GDS.JRTCKXETXF": {
					"effectiveDate": "2020-06-01T00:00:00Z",
					"offerTermCode": "JRTCKXETXF",
					"priceDimensions": {
						"XUZVGXG9M6A44GDS.JRTCKXETXF.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "$0.17 per Unused Reservation Linux c5.xlarge Instance Hour",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.1700000000"
							},
							"rateCode": "XUZVGXG9M6A44GDS.JRTCKXETXF.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "XUZVGXG9M6A44GDS",
					"termAttributes": {}
				}
			}
		},
		"version": "20200618221809"
	},
	{
		"product": {
			"attributes": {
				"capacitystatus": "AllocatedCapacityReservation",
				"clockSpeed": "3 GHz",
				"currentGeneration": "Yes",
				"dedicatedEbsThroughput": "Up to 2250 Mbps",
				"ecu": "20",
				"enhancedNetworkingSupported": "Yes",
				"instanceFamily": "Compute optimized",
				"instanceType": "c5.xlarge",
				"instancesku": "35AEEWH98DECPC35",
				"intelAvx2Available": "Yes",
				"intelAvxAvailable": "Yes",
				"intelTurboAvailable": "Yes",
				"licenseModel": "No License required",
				"location": "US East (N. Virginia)",
				"locationType": "AWS Region",
				"memory": "8 GiB",
				"networkPerformance": "Up to 10 Gigabit",
				"normalizationSizeFactor": "8",
				"operatingSystem": "Linux",
				"operation": "RunInstances",
				"physicalProcessor": "Intel Xeon Platinum 8124M",
				"preInstalledSw": "NA",
				"processorArchitecture": "64-bit",
				"processorFeatures": "Intel AVX; Intel AVX2; Intel AVX512; Intel Turbo",
				"servicecode": "AmazonEC2",
				"servicename": "Amazon Elastic Compute Cloud",
				"storage": "EBS only",
				"tenancy": "Shared",
				"usagetype": "Reservation:c5.xlarge",
				"vcpu": "4"
			},
			"productFamily": "Compute Instance",
			"sku": "Z5FBXQS9D7E49Y7E"
		},
		"publicationDate": "2020-06-18T22:18:09Z",
		"serviceCode": "AmazonEC2",
		"terms": {
			"OnDemand": {
				"Z5FBXQS9D7E49Y7E.JRTCKXETXF": {
					"effectiveDate": "2020-06-01T00:00:00Z",
					"offerTermCode": "JRTCKXETXF",
					"priceDimensions": {
						"Z5FBXQS9D7E49Y7E.JRTCKXETXF.6YS6EN2CT7": {
							"appliesTo": [],
							"beginRange": "0",
							"description": "$0.00 per Reservation Linux c5.xlarge Instance Hour",
							"endRange": "Inf",
							"pricePerUnit": {
								"USD": "0.0000000000"
							},
							"rateCode": "Z5FBXQS9D7E49Y7E.JRTCKXETXF.6YS6EN2CT7",
							"unit": "Hrs"
						}
					},
					"sku": "Z5FBXQS9D7E49Y7E",
					"termAttributes": {}
				}
			}
		},
		"version": "20200618221809"
	}
]
`)
)
