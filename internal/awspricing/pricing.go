package awspricing

// type PricingResponse struct {
// 	PriceList []PriceOffering `json:"priceList"`
// }

// type PriceOffering struct {
// 	Product         Product   `json:"product"`
// 	PublicationDate time.Time `json:"publicationDate"`
// 	ServiceCode     string    `json:"serviceCode"`
// 	Terms           Terms     `json:"terms,omitempty"`
// 	Version         string    `json:"version"`
// 	Terms           Terms     `json:"terms,omitempty"`
// 	Terms           Terms     `json:"terms,omitempty"`
// }

// type Attributes struct {
// 	Capacitystatus              string `json:"capacitystatus"`
// 	ClockSpeed                  string `json:"clockSpeed"`
// 	CurrentGeneration           string `json:"currentGeneration"`
// 	DedicatedEbsThroughput      string `json:"dedicatedEbsThroughput"`
// 	Ecu                         string `json:"ecu"`
// 	EnhancedNetworkingSupported string `json:"enhancedNetworkingSupported"`
// 	InstanceFamily              string `json:"instanceFamily"`
// 	InstanceType                string `json:"instanceType"`
// 	IntelAvx2Available          string `json:"intelAvx2Available"`
// 	IntelAvxAvailable           string `json:"intelAvxAvailable"`
// 	IntelTurboAvailable         string `json:"intelTurboAvailable"`
// 	LicenseModel                string `json:"licenseModel"`
// 	Location                    string `json:"location"`
// 	LocationType                string `json:"locationType"`
// 	Memory                      string `json:"memory"`
// 	NetworkPerformance          string `json:"networkPerformance"`
// 	NormalizationSizeFactor     string `json:"normalizationSizeFactor"`
// 	OperatingSystem             string `json:"operatingSystem"`
// 	Operation                   string `json:"operation"`
// 	PhysicalProcessor           string `json:"physicalProcessor"`
// 	PreInstalledSw              string `json:"preInstalledSw"`
// 	ProcessorArchitecture       string `json:"processorArchitecture"`
// 	ProcessorFeatures           string `json:"processorFeatures"`
// 	Servicecode                 string `json:"servicecode"`
// 	Servicename                 string `json:"servicename"`
// 	Storage                     string `json:"storage"`
// 	Tenancy                     string `json:"tenancy"`
// 	Usagetype                   string `json:"usagetype"`
// 	Vcpu                        string `json:"vcpu"`
// }
// type Product struct {
// 	Attributes    Attributes `json:"attributes"`
// 	ProductFamily string     `json:"productFamily"`
// 	Sku           string     `json:"sku"`
// }
// type PricePerUnit struct {
// 	USD string `json:"USD"`
// }

// type Three5AEEWH98DECPC35JRTCKXETXF6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type PriceDimensions struct {
// 	Three5AEEWH98DECPC35JRTCKXETXF6YS6EN2CT7 Three5AEEWH98DECPC35JRTCKXETXF6YS6EN2CT7 `json:"35AEEWH98DECPC35.JRTCKXETXF.6YS6EN2CT7"`
// }
// type TermAttributes struct {
// }
// type Three5AEEWH98DECPC35JRTCKXETXF struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type OnDemand struct {
// 	Three5AEEWH98DECPC35JRTCKXETXF Three5AEEWH98DECPC35JRTCKXETXF `json:"35AEEWH98DECPC35.JRTCKXETXF"`
// }
// type Three5AEEWH98DECPC3538NPMPTW362TG2D8R56U struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	Description  string        `json:"description"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type Three5AEEWH98DECPC3538NPMPTW366YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// // type PriceDimensions struct {
// // 	Three5AEEWH98DECPC3538NPMPTW362TG2D8R56U Three5AEEWH98DECPC3538NPMPTW362TG2D8R56U `json:"35AEEWH98DECPC35.38NPMPTW36.2TG2D8R56U"`
// // 	Three5AEEWH98DECPC3538NPMPTW366YS6EN2CT7 Three5AEEWH98DECPC3538NPMPTW366YS6EN2CT7 `json:"35AEEWH98DECPC35.38NPMPTW36.6YS6EN2CT7"`
// // }
// type TermAttributes struct {
// 	LeaseContractLength string `json:"LeaseContractLength"`
// 	OfferingClass       string `json:"OfferingClass"`
// 	PurchaseOption      string `json:"PurchaseOption"`
// }
// type Three5AEEWH98DECPC3538NPMPTW36 struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC354NA7Y494T46YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// // type PriceDimensions struct {
// // 	Three5AEEWH98DECPC354NA7Y494T46YS6EN2CT7 Three5AEEWH98DECPC354NA7Y494T46YS6EN2CT7 `json:"35AEEWH98DECPC35.4NA7Y494T4.6YS6EN2CT7"`
// // }
// type Three5AEEWH98DECPC354NA7Y494T4 struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC356QCMYABX3D2TG2D8R56U struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	Description  string        `json:"description"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type Three5AEEWH98DECPC356QCMYABX3D6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// // type PriceDimensions struct {
// // 	Three5AEEWH98DECPC356QCMYABX3D2TG2D8R56U Three5AEEWH98DECPC356QCMYABX3D2TG2D8R56U `json:"35AEEWH98DECPC35.6QCMYABX3D.2TG2D8R56U"`
// // 	Three5AEEWH98DECPC356QCMYABX3D6YS6EN2CT7 Three5AEEWH98DECPC356QCMYABX3D6YS6EN2CT7 `json:"35AEEWH98DECPC35.6QCMYABX3D.6YS6EN2CT7"`
// // }
// type Three5AEEWH98DECPC356QCMYABX3D struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC357NE97W5U4E6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// // type PriceDimensions struct {
// // 	Three5AEEWH98DECPC357NE97W5U4E6YS6EN2CT7 Three5AEEWH98DECPC357NE97W5U4E6YS6EN2CT7 `json:"35AEEWH98DECPC35.7NE97W5U4E.6YS6EN2CT7"`
// // }
// type Three5AEEWH98DECPC357NE97W5U4E struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC35BPH4J8HBKS6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// // type PriceDimensions struct {
// // 	Three5AEEWH98DECPC35BPH4J8HBKS6YS6EN2CT7 Three5AEEWH98DECPC35BPH4J8HBKS6YS6EN2CT7 `json:"35AEEWH98DECPC35.BPH4J8HBKS.6YS6EN2CT7"`
// // }
// type Three5AEEWH98DECPC35BPH4J8HBKS struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC35CUZHX8X6JH2TG2D8R56U struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	Description  string        `json:"description"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type Three5AEEWH98DECPC35CUZHX8X6JH6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// // type PriceDimensions struct {
// // 	Three5AEEWH98DECPC35CUZHX8X6JH2TG2D8R56U Three5AEEWH98DECPC35CUZHX8X6JH2TG2D8R56U `json:"35AEEWH98DECPC35.CUZHX8X6JH.2TG2D8R56U"`
// // 	Three5AEEWH98DECPC35CUZHX8X6JH6YS6EN2CT7 Three5AEEWH98DECPC35CUZHX8X6JH6YS6EN2CT7 `json:"35AEEWH98DECPC35.CUZHX8X6JH.6YS6EN2CT7"`
// // }
// type Three5AEEWH98DECPC35CUZHX8X6JH struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC35HU7G6KETJZ2TG2D8R56U struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	Description  string        `json:"description"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type Three5AEEWH98DECPC35HU7G6KETJZ6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// // type PriceDimensions struct {
// // 	Three5AEEWH98DECPC35HU7G6KETJZ2TG2D8R56U Three5AEEWH98DECPC35HU7G6KETJZ2TG2D8R56U `json:"35AEEWH98DECPC35.HU7G6KETJZ.2TG2D8R56U"`
// // 	Three5AEEWH98DECPC35HU7G6KETJZ6YS6EN2CT7 Three5AEEWH98DECPC35HU7G6KETJZ6YS6EN2CT7 `json:"35AEEWH98DECPC35.HU7G6KETJZ.6YS6EN2CT7"`
// // }
// type Three5AEEWH98DECPC35HU7G6KETJZ struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC35MZU6U2429S2TG2D8R56U struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	Description  string        `json:"description"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type Three5AEEWH98DECPC35MZU6U2429S6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// // type PriceDimensions struct {
// // 	Three5AEEWH98DECPC35MZU6U2429S2TG2D8R56U Three5AEEWH98DECPC35MZU6U2429S2TG2D8R56U `json:"35AEEWH98DECPC35.MZU6U2429S.2TG2D8R56U"`
// // 	Three5AEEWH98DECPC35MZU6U2429S6YS6EN2CT7 Three5AEEWH98DECPC35MZU6U2429S6YS6EN2CT7 `json:"35AEEWH98DECPC35.MZU6U2429S.6YS6EN2CT7"`
// // }
// type Three5AEEWH98DECPC35MZU6U2429S struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC35NQ3QZPMQV92TG2D8R56U struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	Description  string        `json:"description"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type Three5AEEWH98DECPC35NQ3QZPMQV96YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type PriceDimensions struct {
// 	Three5AEEWH98DECPC35NQ3QZPMQV92TG2D8R56U Three5AEEWH98DECPC35NQ3QZPMQV92TG2D8R56U `json:"35AEEWH98DECPC35.NQ3QZPMQV9.2TG2D8R56U"`
// 	Three5AEEWH98DECPC35NQ3QZPMQV96YS6EN2CT7 Three5AEEWH98DECPC35NQ3QZPMQV96YS6EN2CT7 `json:"35AEEWH98DECPC35.NQ3QZPMQV9.6YS6EN2CT7"`
// }
// type Three5AEEWH98DECPC35NQ3QZPMQV9 struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC35R5XV2EPZQZ2TG2D8R56U struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	Description  string        `json:"description"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type Three5AEEWH98DECPC35R5XV2EPZQZ6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type PriceDimensions struct {
// 	Three5AEEWH98DECPC35R5XV2EPZQZ2TG2D8R56U Three5AEEWH98DECPC35R5XV2EPZQZ2TG2D8R56U `json:"35AEEWH98DECPC35.R5XV2EPZQZ.2TG2D8R56U"`
// 	Three5AEEWH98DECPC35R5XV2EPZQZ6YS6EN2CT7 Three5AEEWH98DECPC35R5XV2EPZQZ6YS6EN2CT7 `json:"35AEEWH98DECPC35.R5XV2EPZQZ.6YS6EN2CT7"`
// }
// type Three5AEEWH98DECPC35R5XV2EPZQZ struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC35VJWZNREJX22TG2D8R56U struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	Description  string        `json:"description"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type Three5AEEWH98DECPC35VJWZNREJX26YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type PriceDimensions struct {
// 	Three5AEEWH98DECPC35VJWZNREJX22TG2D8R56U Three5AEEWH98DECPC35VJWZNREJX22TG2D8R56U `json:"35AEEWH98DECPC35.VJWZNREJX2.2TG2D8R56U"`
// 	Three5AEEWH98DECPC35VJWZNREJX26YS6EN2CT7 Three5AEEWH98DECPC35VJWZNREJX26YS6EN2CT7 `json:"35AEEWH98DECPC35.VJWZNREJX2.6YS6EN2CT7"`
// }
// type Three5AEEWH98DECPC35VJWZNREJX2 struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Three5AEEWH98DECPC35Z2E3P23VKM6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type PriceDimensions struct {
// 	Three5AEEWH98DECPC35Z2E3P23VKM6YS6EN2CT7 Three5AEEWH98DECPC35Z2E3P23VKM6YS6EN2CT7 `json:"35AEEWH98DECPC35.Z2E3P23VKM.6YS6EN2CT7"`
// }
// type Three5AEEWH98DECPC35Z2E3P23VKM struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type Reserved struct {
// 	Three5AEEWH98DECPC3538NPMPTW36 Three5AEEWH98DECPC3538NPMPTW36 `json:"35AEEWH98DECPC35.38NPMPTW36"`
// 	Three5AEEWH98DECPC354NA7Y494T4 Three5AEEWH98DECPC354NA7Y494T4 `json:"35AEEWH98DECPC35.4NA7Y494T4"`
// 	Three5AEEWH98DECPC356QCMYABX3D Three5AEEWH98DECPC356QCMYABX3D `json:"35AEEWH98DECPC35.6QCMYABX3D"`
// 	Three5AEEWH98DECPC357NE97W5U4E Three5AEEWH98DECPC357NE97W5U4E `json:"35AEEWH98DECPC35.7NE97W5U4E"`
// 	Three5AEEWH98DECPC35BPH4J8HBKS Three5AEEWH98DECPC35BPH4J8HBKS `json:"35AEEWH98DECPC35.BPH4J8HBKS"`
// 	Three5AEEWH98DECPC35CUZHX8X6JH Three5AEEWH98DECPC35CUZHX8X6JH `json:"35AEEWH98DECPC35.CUZHX8X6JH"`
// 	Three5AEEWH98DECPC35HU7G6KETJZ Three5AEEWH98DECPC35HU7G6KETJZ `json:"35AEEWH98DECPC35.HU7G6KETJZ"`
// 	Three5AEEWH98DECPC35MZU6U2429S Three5AEEWH98DECPC35MZU6U2429S `json:"35AEEWH98DECPC35.MZU6U2429S"`
// 	Three5AEEWH98DECPC35NQ3QZPMQV9 Three5AEEWH98DECPC35NQ3QZPMQV9 `json:"35AEEWH98DECPC35.NQ3QZPMQV9"`
// 	Three5AEEWH98DECPC35R5XV2EPZQZ Three5AEEWH98DECPC35R5XV2EPZQZ `json:"35AEEWH98DECPC35.R5XV2EPZQZ"`
// 	Three5AEEWH98DECPC35VJWZNREJX2 Three5AEEWH98DECPC35VJWZNREJX2 `json:"35AEEWH98DECPC35.VJWZNREJX2"`
// 	Three5AEEWH98DECPC35Z2E3P23VKM Three5AEEWH98DECPC35Z2E3P23VKM `json:"35AEEWH98DECPC35.Z2E3P23VKM"`
// }
// type Terms struct {
// 	OnDemand OnDemand `json:"OnDemand"`
// 	Reserved Reserved `json:"Reserved"`
// }
// type XUZVGXG9M6A44GDSJRTCKXETXF6YS6EN2CT7 struct {
// 	AppliesTo    []interface{} `json:"appliesTo"`
// 	BeginRange   string        `json:"beginRange"`
// 	Description  string        `json:"description"`
// 	EndRange     string        `json:"endRange"`
// 	PricePerUnit PricePerUnit  `json:"pricePerUnit"`
// 	RateCode     string        `json:"rateCode"`
// 	Unit         string        `json:"unit"`
// }
// type PriceDimensions struct {
// 	XUZVGXG9M6A44GDSJRTCKXETXF6YS6EN2CT7 XUZVGXG9M6A44GDSJRTCKXETXF6YS6EN2CT7 `json:"XUZVGXG9M6A44GDS.JRTCKXETXF.6YS6EN2CT7"`
// }
// type TermAttributes struct {
// }
// type XUZVGXG9M6A44GDSJRTCKXETXF struct {
// 	EffectiveDate   time.Time       `json:"effectiveDate"`
// 	OfferTermCode   string          `json:"offerTermCode"`
// 	PriceDimensions PriceDimensions `json:"priceDimensions"`
// 	Sku             string          `json:"sku"`
// 	TermAttributes  TermAttributes  `json:"termAttributes"`
// }
// type OnDemand struct {
// 	XUZVGXG9M6A44GDSJRTCKXETXF XUZVGXG9M6A44GDSJRTCKXETXF `json:"XUZVGXG9M6A44GDS.JRTCKXETXF"`
// }
// type Terms struct {
// 	OnDemand OnDemand `json:"OnDemand"`
// }

// onDemand := resp.PriceList[0]["terms"].(map[string]interface{})["OnDemand"]

// keys := reflect.ValueOf(onDemand).MapKeys()

// productCode := keys[0]

// priceDimensions := onDemand.(map[string]interface{})[productCode.String()].(map[string]interface{})["priceDimensions"]

// pcKeys := reflect.ValueOf(priceDimensions).MapKeys()

// priceDimensionsKey := pcKeys[0]

// price := priceDimensions.(map[string]interface{})[priceDimensionsKey.String()].(map[string]interface{})["pricePerUnit"].(map[string]interface{})["USD"].(string)
// priceUnit := priceDimensions.(map[string]interface{})[priceDimensionsKey.String()].(map[string]interface{})["unit"].(string)

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
