package world

//go:generate reform

// Countrylanguage represents a row in countrylanguage table.
//reform:countrylanguage
type Countrylanguage struct {
	CountryCode string `reform:"CountryCode"`
	Language    string `reform:"Language"`
	IsOfficial  []byte `reform:"IsOfficial"` // FIXME unhandled database type "enum"
	Percentage  string `reform:"Percentage"`
}
