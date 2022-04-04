package pagila

import (
	"time"
)

//go:generate reform

// FilmActor represents a row in film_actor table.
//reform:film_actor
type FilmActor struct {
	ActorID    int32     `reform:"actor_id"`
	FilmID     int32     `reform:"film_id"`
	LastUpdate time.Time `reform:"last_update"`
}
