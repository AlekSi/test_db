package sakila

import (
	"time"
)

//go:generate reform

// Staff represents a row in staff table.
//reform:staff
type Staff struct {
	StaffID    int8      `reform:"staff_id,pk"`
	FirstName  string    `reform:"first_name"`
	LastName   string    `reform:"last_name"`
	AddressID  int16     `reform:"address_id"`
	Picture    []byte    `reform:"picture"`
	Email      *string   `reform:"email"`
	StoreID    int8      `reform:"store_id"`
	Active     int8      `reform:"active"`
	Username   string    `reform:"username"`
	Password   *string   `reform:"password"`
	LastUpdate time.Time `reform:"last_update"`
}
