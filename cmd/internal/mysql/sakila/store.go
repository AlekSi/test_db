package sakila

import (
	"time"
)

//go:generate reform

// Store represents a row in store table.
//reform:store
type Store struct {
	StoreID        int8      `reform:"store_id,pk"`
	ManagerStaffID int8      `reform:"manager_staff_id"`
	AddressID      int16     `reform:"address_id"`
	LastUpdate     time.Time `reform:"last_update"`
}
