package pagila

//go:generate reform

// CustomerList represents a row in customer_list table.
//reform:customer_list
type CustomerList struct {
	ID      *int32  `reform:"id"`
	Name    *string `reform:"name"`
	Address *string `reform:"address"`
	ZipCode *string `reform:"zip code"`
	Phone   *string `reform:"phone"`
	City    *string `reform:"city"`
	Country *string `reform:"country"`
	Notes   *string `reform:"notes"`
	Sid     *int32  `reform:"sid"`
}
