package sakila

//go:generate reform

// CustomerList represents a row in customer_list table.
//reform:customer_list
type CustomerList struct {
	ID      int16   `reform:"ID"`
	Name    *string `reform:"name"`
	Address string  `reform:"address"`
	ZipCode *string `reform:"zip code"`
	Phone   string  `reform:"phone"`
	City    string  `reform:"city"`
	Country string  `reform:"country"`
	Notes   string  `reform:"notes"`
	SID     int8    `reform:"SID"`
}
