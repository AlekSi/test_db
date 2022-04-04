package pagila

import (
	"time"
)

//go:generate reform

// FilmCategory represents a row in film_category table.
//reform:film_category
type FilmCategory struct {
	FilmID     int32     `reform:"film_id"`
	CategoryID int32     `reform:"category_id"`
	LastUpdate time.Time `reform:"last_update"`
}
