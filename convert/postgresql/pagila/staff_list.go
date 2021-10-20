package pagila

//go:generate reform

// StaffList represents a row in staff_list table.
//reform:staff_list
type StaffList struct {
	ID      *int32  `reform:"id"`
	Name    *string `reform:"name"`
	Address *string `reform:"address"`
	ZipCode *string `reform:"zip code"`
	Phone   *string `reform:"phone"`
	City    *string `reform:"city"`
	Country *string `reform:"country"`
	Sid     *int32  `reform:"sid"`
}
