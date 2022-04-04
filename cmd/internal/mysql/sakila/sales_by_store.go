package sakila

//go:generate reform

// SalesByStore represents a row in sales_by_store table.
//reform:sales_by_store
type SalesByStore struct {
	Store      *string `reform:"store"`
	Manager    *string `reform:"manager"`
	TotalSales *string `reform:"total_sales"`
}
