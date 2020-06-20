package model

import (
	"movies/src/domain/exception"
	"movies/src/domain/validator"
)

type Movie struct {
	Id        string
	User_id   int64
	Title     string
	Length    int
	Synopsis  string
	Image_url string
}

func (m *Movie) Create(
	User_id int64,
	Title string,
	Length int,
	Synopsis string,
	Image_url string) (Movie, error) {
	if err := validator.ValidateRequired(Title, "movie title required"); err != nil {
		return Movie{}, exception.InvalidMovieName{ErrMessage: err.Error()}
	}
	if err := validator.ValidateLength(Length, "movie length invalid"); err != nil {
		return Movie{}, exception.InvalidMovieLength{ErrMessage: err.Error()}
	}

	return Movie{
			User_id:   User_id,
			Title:     Title,
			Length:    Length,
			Synopsis:  Synopsis,
			Image_url: Image_url},
		nil

}
