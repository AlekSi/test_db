package pagila

//go:generate reform

// FilmList represents a row in film_list table.
//reform:film_list
type FilmList struct {
	Fid         *int32  `reform:"fid"`
	Title       *string `reform:"title"`
	Description *string `reform:"description"`
	Category    *string `reform:"category"`
	Price       *string `reform:"price"`
	Length      *int16  `reform:"length"`
	Rating      []byte  `reform:"rating"` // FIXME unhandled database type "USER-DEFINED"
	Actors      *string `reform:"actors"`
}
