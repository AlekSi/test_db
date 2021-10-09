package sakila

import (
	"time"
)

//go:generate reform

// Language represents a row in language table.
//reform:language
type Language struct {
	LanguageID int8      `reform:"language_id,pk"`
	Name       string    `reform:"name"`
	LastUpdate time.Time `reform:"last_update"`
}
