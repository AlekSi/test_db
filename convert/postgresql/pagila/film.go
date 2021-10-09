package pagila

import (
	"time"
)

//go:generate reform

// Film represents a row in film table.
//reform:film
type Film struct {
	FilmID             int32     `reform:"film_id,pk"`
	Title              string    `reform:"title"`
	Description        *string   `reform:"description"`
	ReleaseYear        *int32    `reform:"release_year"`
	LanguageID         int32     `reform:"language_id"`
	OriginalLanguageID *int32    `reform:"original_language_id"`
	RentalDuration     int16     `reform:"rental_duration"`
	RentalRate         string    `reform:"rental_rate"`
	Length             *int16    `reform:"length"`
	ReplacementCost    string    `reform:"replacement_cost"`
	Rating             []byte    `reform:"rating"` // FIXME unhandled database type "USER-DEFINED"
	LastUpdate         time.Time `reform:"last_update"`
	SpecialFeatures    []byte    `reform:"special_features"` // FIXME unhandled database type "ARRAY"
	Fulltext           []byte    `reform:"fulltext"`         // FIXME unhandled database type "tsvector"
}
