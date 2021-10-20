package sakila

import (
	"time"
)

//go:generate reform

// FilmCategory represents a row in film_category table.
//reform:film_category
type FilmCategory struct {
	FilmID     int16     `reform:"film_id"`
	CategoryID int8      `reform:"category_id"`
	LastUpdate time.Time `reform:"last_update"`
}
