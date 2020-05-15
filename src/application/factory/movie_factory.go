package factory

import (
	"movies/src/application/command"
	"movies/src/domain/model"
)

func CreateMovie(command command.MovieCommand) (model.Movie, error) {
	var movie model.Movie
	movie, err := movie.Create(command.UserId, command.Title, command.Length, command.Synopsis, command.ImageUrl)
	return movie, err
}
