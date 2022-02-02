package world

//go:generate reform

// City represents a row in city table.
//reform:city
type City struct {
	ID          int32  `reform:"ID,pk"`
	Name        string `reform:"Name"`
	CountryCode string `reform:"CountryCode"`
	District    string `reform:"District"`
	Population  int32  `reform:"Population"`
}
