package factory

import (
	"movies/src/application/command"
	"movies/src/domain/model"
)

func CreateMovie(command command.MovieCommand) (model.Movie, error) {
	var movie model.Movie
	movie, err := movie.Create(command.User_id, command.Title, command.Length, command.Synopsis, command.Image_url)
	return movie, err
}
