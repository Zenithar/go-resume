package schema

// Location represents location information
type Location struct {
	CountryCode string `yaml:"countryCode"`
	Region      string `yaml:"region"`
	PostalCode  int64  `yaml:"postalCode"`
	City        string `yaml:"city"`
	Address     string `yaml:"address"`
}
