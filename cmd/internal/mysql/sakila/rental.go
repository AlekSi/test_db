package sakila

import (
	"time"
)

//go:generate reform

// Rental represents a row in rental table.
//reform:rental
type Rental struct {
	RentalID    int32      `reform:"rental_id,pk"`
	RentalDate  time.Time  `reform:"rental_date"`
	InventoryID int32      `reform:"inventory_id"`
	CustomerID  int16      `reform:"customer_id"`
	ReturnDate  *time.Time `reform:"return_date"`
	StaffID     int8       `reform:"staff_id"`
	LastUpdate  time.Time  `reform:"last_update"`
}
