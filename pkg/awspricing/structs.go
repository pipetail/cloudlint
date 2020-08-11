package awspricing

import "time"

// GetProductResponse is a response struct from pricing api
type GetProductResponse struct {
	FormatVersion string      `json:"FormatVersion"`
	NextToken     interface{} `json:"NextToken"`
	PriceList     []PriceList `json:"PriceList"`
}

// PriceList contains product metadata and price information
type PriceList struct {
	Product         Product   `json:"product"`
	PublicationDate time.Time `json:"publicationDate"`
	ServiceCode     string    `json:"serviceCode"`
	Terms           Terms     `json:"terms"`
	Version         string    `json:"version"`
}

// Attributes for EC2 instance products
type Attributes struct {
	Capacitystatus              string `json:"capacitystatus"`
	ClockSpeed                  string `json:"clockSpeed"`
	CurrentGeneration           string `json:"currentGeneration"`
	DedicatedEbsThroughput      string `json:"dedicatedEbsThroughput"`
	Ecu                         string `json:"ecu"`
	EnhancedNetworkingSupported string `json:"enhancedNetworkingSupported"`
	InstanceFamily              string `json:"instanceFamily"`
	InstanceType                string `json:"instanceType"`
	IntelAvx2Available          string `json:"intelAvx2Available"`
	IntelAvxAvailable           string `json:"intelAvxAvailable"`
	IntelTurboAvailable         string `json:"intelTurboAvailable"`
	LicenseModel                string `json:"licenseModel"`
	Location                    string `json:"location"`
	LocationType                string `json:"locationType"`
	Memory                      string `json:"memory"`
	NetworkPerformance          string `json:"networkPerformance"`
	NormalizationSizeFactor     string `json:"normalizationSizeFactor"`
	OperatingSystem             string `json:"operatingSystem"`
	Operation                   string `json:"operation"`
	PhysicalProcessor           string `json:"physicalProcessor"`
	PreInstalledSw              string `json:"preInstalledSw"`
	ProcessorArchitecture       string `json:"processorArchitecture"`
	ProcessorFeatures           string `json:"processorFeatures"`
	Servicecode                 string `json:"servicecode"`
	Servicename                 string `json:"servicename"`
	Storage                     string `json:"storage"`
	Tenancy                     string `json:"tenancy"`
	Usagetype                   string `json:"usagetype"`
	Vcpu                        string `json:"vcpu"`
}

// Product holds more detailed info about the concrete product
type Product struct {
	Attributes    Attributes `json:"attributes"`
	ProductFamily string     `json:"productFamily"`
	Sku           string     `json:"sku"`
}

// PricePerUnit holds price per unit
type PricePerUnit struct {
	USD string `json:"USD"`
}

// PriceDimension holds more metadata about product prices
type PriceDimension struct {
	AppliesTo    []interface{} `json:"appliesTo"`
	BeginRange   string        `json:"beginRange"`
	Description  string        `json:"description"`
	EndRange     string        `json:"endRange"`
	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
	RateCode     string        `json:"rateCode"`
	Unit         string        `json:"unit"`
}

// TermAttributes are metadata for pricing terms
type TermAttributes struct {
	LeaseContractLength string `json:"LeaseContractLength"`
	OfferingClass       string `json:"OfferingClass"`
	PurchaseOption      string `json:"PurchaseOption"`
}

// Offering is an OnDemand or Reserved price offer
type Offering struct {
	EffectiveDate   time.Time                 `json:"effectiveDate"`
	OfferTermCode   string                    `json:"offerTermCode"`
	PriceDimensions map[string]PriceDimension `json:"priceDimensions"`
	Sku             string                    `json:"sku"`
	TermAttributes  TermAttributes            `json:"termAttributes"`
}

// Terms for products (OnDemand and Reserved)
type Terms struct {
	OnDemand map[string]Offering `json:"OnDemand"`
	Reserved map[string]Offering `json:"Reserved"`
}
