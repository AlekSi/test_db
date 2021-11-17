package world

//go:generate reform

// CountryLanguage represents a row in countrylanguage table.
//reform:countrylanguage
type CountryLanguage struct {
	CountryCode string `reform:"CountryCode"`
	Language    string `reform:"Language"`
	IsOfficial  []byte `reform:"IsOfficial"` // FIXME unhandled database type "enum"
	Percentage  string `reform:"Percentage"`
}
