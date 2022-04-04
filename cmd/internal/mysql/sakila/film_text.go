package sakila

//go:generate reform

// FilmText represents a row in film_text table.
//reform:film_text
type FilmText struct {
	FilmID      int16   `reform:"film_id,pk"`
	Title       string  `reform:"title"`
	Description *string `reform:"description"`
}
