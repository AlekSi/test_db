package sakila

import (
	"time"
)

//go:generate reform

// Payment represents a row in payment table.
//reform:payment
type Payment struct {
	PaymentID   int16      `reform:"payment_id,pk"`
	CustomerID  int16      `reform:"customer_id"`
	StaffID     int8       `reform:"staff_id"`
	RentalID    *int32     `reform:"rental_id"`
	Amount      string     `reform:"amount"`
	PaymentDate time.Time  `reform:"payment_date"`
	LastUpdate  *time.Time `reform:"last_update"`
}
