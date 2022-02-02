package sakila

//go:generate reform

// ActorInfo represents a row in actor_info table.
//reform:actor_info
type ActorInfo struct {
	ActorID   int16   `reform:"actor_id"`
	FirstName string  `reform:"first_name"`
	LastName  string  `reform:"last_name"`
	FilmInfo  *string `reform:"film_info"`
}
