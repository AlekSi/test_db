package sakila

import (
	"time"
)

//go:generate reform

// Customer represents a row in customer table.
//reform:customer
type Customer struct {
	CustomerID int16      `reform:"customer_id,pk"`
	StoreID    int8       `reform:"store_id"`
	FirstName  string     `reform:"first_name"`
	LastName   string     `reform:"last_name"`
	Email      *string    `reform:"email"`
	AddressID  int16      `reform:"address_id"`
	Active     int8       `reform:"active"`
	CreateDate time.Time  `reform:"create_date"`
	LastUpdate *time.Time `reform:"last_update"`
}
