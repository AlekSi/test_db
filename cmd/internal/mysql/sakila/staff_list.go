package sakila

//go:generate reform

// StaffList represents a row in staff_list table.
//reform:staff_list
type StaffList struct {
	ID      int8    `reform:"ID"`
	Name    *string `reform:"name"`
	Address string  `reform:"address"`
	ZipCode *string `reform:"zip code"`
	Phone   string  `reform:"phone"`
	City    string  `reform:"city"`
	Country string  `reform:"country"`
	SID     int8    `reform:"SID"`
}
