package world

//go:generate reform

// City represents a row in city table.
//reform:city
type City struct {
	ID          int32  `reform:"id,pk"`
	Name        string `reform:"name"`
	CountryCode string `reform:"countrycode"`
	District    string `reform:"district"`
	Population  int32  `reform:"population"`
}
