package sakila

import (
	"time"
)

//go:generate reform

// City represents a row in city table.
//reform:city
type City struct {
	CityID     int16     `reform:"city_id,pk"`
	City       string    `reform:"city"`
	CountryID  int16     `reform:"country_id"`
	LastUpdate time.Time `reform:"last_update"`
}
