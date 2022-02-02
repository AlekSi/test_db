package pagila

import (
	"time"
)

//go:generate reform

// Staff represents a row in staff table.
//reform:staff
type Staff struct {
	StaffID    int32     `reform:"staff_id,pk"`
	FirstName  string    `reform:"first_name"`
	LastName   string    `reform:"last_name"`
	AddressID  int32     `reform:"address_id"`
	Email      *string   `reform:"email"`
	StoreID    int32     `reform:"store_id"`
	Active     bool      `reform:"active"`
	Username   string    `reform:"username"`
	Password   *string   `reform:"password"`
	LastUpdate time.Time `reform:"last_update"`
	Picture    []byte    `reform:"picture"`
}
