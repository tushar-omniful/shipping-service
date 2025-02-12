package constants

const (
	AddressTypeCoordinates = "COORDINATES"
	AddressTypeFreeText    = "FREE_TEXT"
	AddressTypeLookup      = "LOOKUP"
)

// Contact Types
const (
	ContactTypeEnterprise = "ENTERPRISE"
	ContactTypeIndividual = "INDIVIDUAL"
)

// Service Codes
const (
	ServiceCodeCOD = "COD"
)

// CountriesWithoutPostalCode defines countries that don't use postal codes
var CountriesWithoutPostalCode = map[string]bool{
	"AE": true,
	"SA": true,
}
