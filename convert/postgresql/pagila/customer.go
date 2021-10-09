package pagila

import (
	"time"
)

//go:generate reform

// Customer represents a row in customer table.
//reform:customer
type Customer struct {
	CustomerID int32      `reform:"customer_id,pk"`
	StoreID    int32      `reform:"store_id"`
	FirstName  string     `reform:"first_name"`
	LastName   string     `reform:"last_name"`
	Email      *string    `reform:"email"`
	AddressID  int32      `reform:"address_id"`
	Activebool bool       `reform:"activebool"`
	CreateDate time.Time  `reform:"create_date"`
	LastUpdate *time.Time `reform:"last_update"`
	Active     *int32     `reform:"active"`
}
