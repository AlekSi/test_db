package pagila

import (
	"time"
)

//go:generate reform

// PaymentP202003 represents a row in payment_p2020_03 table.
//reform:payment_p2020_03
type PaymentP202003 struct {
	PaymentID   int32     `reform:"payment_id"`
	CustomerID  int32     `reform:"customer_id"`
	StaffID     int32     `reform:"staff_id"`
	RentalID    int32     `reform:"rental_id"`
	Amount      string    `reform:"amount"`
	PaymentDate time.Time `reform:"payment_date"`
}
