package pagila

import (
	"time"
)

//go:generate reform

// Inventory represents a row in inventory table.
//reform:inventory
type Inventory struct {
	InventoryID int32     `reform:"inventory_id,pk"`
	FilmID      int32     `reform:"film_id"`
	StoreID     int32     `reform:"store_id"`
	LastUpdate  time.Time `reform:"last_update"`
}
