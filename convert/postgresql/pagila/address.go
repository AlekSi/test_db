package pagila

import (
	"time"
)

//go:generate reform

// Address represents a row in address table.
//reform:address
type Address struct {
	AddressID  int32     `reform:"address_id,pk"`
	Address    string    `reform:"address"`
	Address2   *string   `reform:"address2"`
	District   string    `reform:"district"`
	CityID     int32     `reform:"city_id"`
	PostalCode *string   `reform:"postal_code"`
	Phone      string    `reform:"phone"`
	LastUpdate time.Time `reform:"last_update"`
}
