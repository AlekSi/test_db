package sakila

//go:generate reform

// FilmList represents a row in film_list table.
//reform:film_list
type FilmList struct {
	FID         *int16  `reform:"FID"`
	Title       *string `reform:"title"`
	Description *string `reform:"description"`
	Category    string  `reform:"category"`
	Price       *string `reform:"price"`
	Length      *int16  `reform:"length"`
	Rating      []byte  `reform:"rating"` // FIXME unhandled database type "enum"
	Actors      *string `reform:"actors"`
}
