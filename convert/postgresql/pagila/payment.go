package pagila

import (
	"time"
)

//go:generate reform

// Payment represents a row in payment table.
//reform:payment
type Payment struct {
	PaymentID   int32     `reform:"payment_id"`
	CustomerID  int32     `reform:"customer_id"`
	StaffID     int32     `reform:"staff_id"`
	RentalID    int32     `reform:"rental_id"`
	Amount      string    `reform:"amount"`
	PaymentDate time.Time `reform:"payment_date"`
}
