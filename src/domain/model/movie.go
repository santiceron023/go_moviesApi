package model

import "movies/src/domain/exception"

type Movie struct {
	Id        string
	User_id   int64
	Title     string
	Length    int
	Synopsis  string
	Image_url string
}

func (m *Movie) Validate() error {
	if len(m.Title) == 0 {
		//TODO IMPL ERR
		return exception.InvalidMovieName{"movie name must not be empty"}
	}
	if m.Length <= 0 {
		//TODO IMPL ERR
		return exception.InvalidMovieDuration{"invalid movie duration"}
	}
	return nil
}
