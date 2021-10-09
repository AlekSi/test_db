package sakila

import (
	"time"
)

//go:generate reform

// Film represents a row in film table.
//reform:film
type Film struct {
	FilmID             int16     `reform:"film_id,pk"`
	Title              string    `reform:"title"`
	Description        *string   `reform:"description"`
	ReleaseYear        *uint16   `reform:"release_year"`
	LanguageID         int8      `reform:"language_id"`
	OriginalLanguageID *int8     `reform:"original_language_id"`
	RentalDuration     int8      `reform:"rental_duration"`
	RentalRate         string    `reform:"rental_rate"`
	Length             *int16    `reform:"length"`
	ReplacementCost    string    `reform:"replacement_cost"`
	Rating             []byte    `reform:"rating"`           // FIXME unhandled database type "enum"
	SpecialFeatures    []byte    `reform:"special_features"` // FIXME unhandled database type "set"
	LastUpdate         time.Time `reform:"last_update"`
}
