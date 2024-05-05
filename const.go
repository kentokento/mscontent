package mscontent

const BaseURI = "https://content.api.bingads.microsoft.com/shopping/v9.1/bmc/"

const (
	AgeGroupAdult   = "adult"
	AgeGroupKids    = "kids"
	AgeGroupToddler = "toddler"
	AgeGroupInfant  = "infant"
	AgeGroupNewborn = "newborn"
)

const (
	AvailabilityInStock    = "in stock"
	AvailabilityOutOfStock = "out of stock"
	AvailabilityPreorder   = "preorder"
)

const (
	ChannelLocal  = "Local"
	ChannelOnline = "Online"
)

const (
	ConditionNew         = "new"
	ConditionRefurbished = "refurbished"
	ConditionUsed        = "used"
)

const (
	GenderFemale = "female"
	GenderMale   = "male"
	GenderUnisex = "unisex"
)

const (
	GtinEAN  = "EAN"  // (European Article Number)
	GtinESBN = "ISBN" // (International Standard Book Number)
	GtinJAN  = "JAN"  // (Japanese Article Number)
	GtinUPC  = "UPC"  // (Universal Product Code)
)

const (
	SizeAU = "AU"
	SizeDE = "DE"
	SizeFR = "FR"
	SizeUK = "UK"
	SizeUS = "US"
)

const (
	SizeTypeRegular    = "regular"
	SizeTypeMaternity  = "maternity"
	SizeTypePetite     = "petite"
	SizeTypePlus       = "plus"
	SizeTypeBigAndTall = "big and tall"
)

const (
	ProductCustomAttributeTypeBoolean       = "boolean"
	ProductCustomAttributeTypeDatetimerange = "datetimerange"
	ProductCustomAttributeTypeFloat         = "float"
	ProductCustomAttributeTypeGroup         = "group"
	ProductCustomAttributeTypeInt           = "int"
	ProductCustomAttributeTypePrice         = "price"
	ProductCustomAttributeTypeText          = "text"
	ProductCustomAttributeTypeTime          = "time"
	ProductCustomAttributeTypeUrl           = "url"
)

const (
	IntentionDefault  = "default"
	IntentionExcluded = "excluded"
	IntentionOptional = "optional"
	IntentionRequired = "required"
)
