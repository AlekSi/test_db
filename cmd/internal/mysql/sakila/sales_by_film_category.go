package sakila

//go:generate reform

// SalesByFilmCategory represents a row in sales_by_film_category table.
//reform:sales_by_film_category
type SalesByFilmCategory struct {
	Category   string  `reform:"category"`
	TotalSales *string `reform:"total_sales"`
}
