package sakila

import (
	"time"
)

//go:generate reform

// Actor represents a row in actor table.
//reform:actor
type Actor struct {
	ActorID    int16     `reform:"actor_id,pk"`
	FirstName  string    `reform:"first_name"`
	LastName   string    `reform:"last_name"`
	LastUpdate time.Time `reform:"last_update"`
}
