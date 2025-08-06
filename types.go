package main

type Project struct {
	ID              string   `json:"id"`
	Title           string   `json:"title"`
	Year            int      `json:"year"`
	Cast            []string `json:"cast,omitempty"`
	Director        string   `json:"director,omitempty"`
	CastingDirector string   `json:"castingDirector,omitempty"`
}

type Actor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Director struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CastingDirector struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CinemaData struct {
	Projects         []Project         `json:"projects"`
	Actors           []Actor           `json:"actors"`
	Directors        []Director        `json:"directors"`
	CastingDirectors []CastingDirector `json:"castingDirectors"`
}
