package sakila

import (
	"time"
)

//go:generate reform

// Country represents a row in country table.
//reform:country
type Country struct {
	CountryID  int16     `reform:"country_id,pk"`
	Country    string    `reform:"country"`
	LastUpdate time.Time `reform:"last_update"`
}
