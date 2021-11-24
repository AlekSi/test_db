package world

//go:generate reform

// Country represents a row in country table.
//reform:country
type Country struct {
	Code           string   `reform:"code,pk"`
	Name           string   `reform:"name"`
	Continent      string   `reform:"continent"`
	Region         string   `reform:"region"`
	SurfaceArea    float32  `reform:"surfacearea"`
	IndepYear      *int16   `reform:"indepyear"`
	Population     int32    `reform:"population"`
	LifeExpectancy *float32 `reform:"lifeexpectancy"`
	GNP            *string  `reform:"gnp"`
	GNPOld         *string  `reform:"gnpold"`
	LocalName      string   `reform:"localname"`
	GovernmentForm string   `reform:"governmentform"`
	HeadOfState    *string  `reform:"headofstate"`
	Capital        *int32   `reform:"capital"`
	Code2          string   `reform:"code2"`
}
