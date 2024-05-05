package mscontent

type Batch struct {
	Entries []Item `json:"entries"`
}

type BatchItemError struct {
	Errors  []Error `json:"errors"`
	Code    string  `json:"code"`
	Message string  `json:"message"`
}

type Error struct {
	Domain       string `json:"domain"`
	Location     string `json:"location"`
	LocationType string `json:"locationType"`
	Message      string `json:"message"`
	Reason       string `json:"reason"`
}

type ErrorResponse struct {
	Error []Errors `json:"error"`
}

type Errors struct {
	Errors   []Error   `json:"errors"`
	Warnings []Warning `json:"warnings"`
	Code     string    `json:"code"`
	Message  string    `json:"message"`
}

type Item struct {
	BatchId    uint32         `json:"batchId"`
	Errors     BatchItemError `json:"errors"`
	MerchantId uint64         `json:"merchantId"`
	Method     string         `json:"method"`
	Product    Product        `json:"product"`
	ProductId  string         `json:"productId"`
}

type Product struct {
	AdditionalImageLinks   []string                 `json:"additionalImageLinks,omitempty"`
	Adult                  bool                     `json:"adult,omitempty"`
	AdwordsGrouping        string                   `json:"adwordsGrouping,omitempty"`
	AdwordsLabels          []string                 `json:"adwordsLabels,omitempty"`
	AdwordsRedirect        string                   `json:"adwordsRedirect,omitempty"`
	AgeGroup               string                   `json:"ageGroup,omitempty"`
	Availability           string                   `json:"availability"`
	AvailabilityDate       string                   `json:"availabilityDate,omitempty"` // ISO 8601 format
	Brand                  string                   `json:"brand"`
	Channel                string                   `json:"channel"`
	Color                  string                   `json:"color"`
	Condition              string                   `json:"condition"`
	ContentLanguage        string                   `json:"contentLanguage"` // ISO 639-1 language code
	CustomAttributes       []ProductCustomAttribute `json:"customAttributes,omitempty"`
	CustomGroups           []ProductCustomGroup     `json:"customGroups,omitempty"`
	CustomLabel0           string                   `json:"customLabel0,omitempty"`
	CustomLabel1           string                   `json:"customLabel1,omitempty"`
	CustomLabel2           string                   `json:"customLabel2,omitempty"`
	CustomLabel3           string                   `json:"customLabel3,omitempty"`
	CustomLabel4           string                   `json:"customLabel4,omitempty"`
	Description            string                   `json:"description,omitempty"`
	Destinations           []ProductDestination     `json:"destinations,omitempty"`
	EnergyEfficiencyClass  string                   `json:"energyEfficiencyClass,omitempty"`
	ExpirationDate         string                   `json:"expirationDate,omitempty"`
	Gender                 string                   `json:"gender,omitempty"`
	GoogleProductCategory  string                   `json:"googleProductCategory"`
	Gtin                   string                   `json:"gtin"`
	ID                     string                   `json:"id,omitempty"`
	IdentifierExists       bool                     `json:"identifierExists"`
	ImageLink              string                   `json:"imageLink"`
	IsBundle               bool                     `json:"isBundle,omitempty"`
	ItemGroupId            string                   `json:"itemGroupId,omitempty"`
	Kind                   string                   `json:"kind,omitempty"`
	Link                   string                   `json:"link"`
	Material               string                   `json:"material,omitempty"`
	MobileLink             string                   `json:"mobileLink,omitempty"`
	Multipack              int                      `json:"multipack,omitempty"`
	Mpn                    string                   `json:"mpn"`
	OfferId                string                   `json:"offerId"`
	OnlineOnly             bool                     `json:"onlineOnly,omitempty"`
	Pattern                string                   `json:"pattern,omitempty"`
	Price                  ProductPrice             `json:"price"`
	ProductType            string                   `json:"productType,omitempty"`
	PromotionId            string                   `json:"promotionId,omitempty"`
	SalePrice              *ProductPrice            `json:"salePrice,omitempty"`
	SalePriceEffectiveDate string                   `json:"salePriceEffectiveDate,omitempty"`
	SellerName             string                   `json:"sellerName,omitempty"`
	Shipping               []ProductShipping        `json:"shipping,omitempty"`       // NOTE: Required if DE.
	ShippingLabel          string                   `json:"shippingLabel,omitempty"`  // NOTE: Required if DE.
	ShippingWeight         *ProductShippingWeight   `json:"shippingWeight,omitempty"` // NOTE: Required if DE.
	Sizes                  []string                 `json:"sizes,omitempty"`
	SizeSystem             string                   `json:"sizeSystem,omitempty"`
	SizeType               string                   `json:"sizeType,omitempty"`
	TargetCountry          string                   `json:"targetCountry"` // ISO 3166 country code
	Taxes                  []ProductTax             `json:"taxes,omitempty"`
	Title                  string                   `json:"title"`
	UnitPricingBaseMeasure *UnitPricing             `json:"unitPricingBaseMeasure,omitempty"`
	UnitPricingMeasure     *UnitPricing             `json:"unitPricingMeasure,omitempty"`
	ValidatedDestinations  []string                 `json:"validatedDestinations,omitempty"`
	Warnings               []Warning                `json:"warnings"`
}

type ProductCustomAttribute struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Unit  string `json:"unit"`
	Value string `json:"value"`
}

type ProductCustomGroup struct {
	Attributes ProductCustomAttribute `json:"attributes"`
	Name       string                 `json:"name"`
}

type ProductDestination struct {
	Intention       string `json:"intention"`
	DestinationName string `json:"destinationName"`
}

type ProductPrice struct {
	Currency string  `json:"currency"` // ISO 4217 currency codes
	Value    float64 `json:"value"`
}

type Products struct {
	Kind          string    `json:"kind"`
	NextPageToken string    `json:"nextPageToken"`
	Resources     []Product `json:"resources"`
}

type ProductShipping struct {
	Country           string       `json:"country"` // ISO 3166 country code
	LocationGroupName string       `json:"locationGroupName"`
	LocationId        string       `json:"locationId"`
	PostalCode        string       `json:"postalCode"`
	Price             ProductPrice `json:"price"`
	Region            string       `json:"region"`
	Service           string       `json:"service"`
}

type ProductShippingWeight struct {
	Unit  string `json:"unit"`
	Value string `json:"value"`
}

type ProductTax struct {
	Country    string  `json:"country"` // ISO 3166 country code
	LocationId string  `json:"locationId"`
	PostalCode string  `json:"postalCode"`
	Rate       float64 `json:"rate"`
	Region     string  `json:"region"`
	TaxShip    bool    `json:"taxShip"`
}

type UnitPricing struct {
	Unit  string  `json:"unit"`
	Value float64 `json:"value"`
}

type Warning struct {
	Domain  string `json:"domain"`
	Message string `json:"message"`
	Reason  string `json:"reason"`
}
