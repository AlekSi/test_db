package world

//go:generate reform

// Country represents a row in country table.
//reform:country
type Country struct {
	Code           string  `reform:"Code,pk"`
	Name           string  `reform:"Name"`
	Continent      []byte  `reform:"Continent"` // FIXME unhandled database type "enum"
	Region         string  `reform:"Region"`
	SurfaceArea    string  `reform:"SurfaceArea"`
	IndepYear      *int16  `reform:"IndepYear"`
	Population     int32   `reform:"Population"`
	LifeExpectancy *string `reform:"LifeExpectancy"`
	GNP            *string `reform:"GNP"`
	GNPOld         *string `reform:"GNPOld"`
	LocalName      string  `reform:"LocalName"`
	GovernmentForm string  `reform:"GovernmentForm"`
	HeadOfState    *string `reform:"HeadOfState"`
	Capital        *int32  `reform:"Capital"`
	Code2          string  `reform:"Code2"`
}
