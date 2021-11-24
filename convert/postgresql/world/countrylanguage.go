package world

//go:generate reform

// CountryLanguage represents a row in countrylanguage table.
//reform:countrylanguage
type CountryLanguage struct {
	CountryCode string  `reform:"countrycode"`
	Language    string  `reform:"language"`
	IsOfficial  bool    `reform:"isofficial"`
	Percentage  float32 `reform:"percentage"`
}
