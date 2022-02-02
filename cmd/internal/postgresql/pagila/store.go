package pagila

import (
	"time"
)

//go:generate reform

// Store represents a row in store table.
//reform:store
type Store struct {
	StoreID        int32     `reform:"store_id,pk"`
	ManagerStaffID int32     `reform:"manager_staff_id"`
	AddressID      int32     `reform:"address_id"`
	LastUpdate     time.Time `reform:"last_update"`
}
