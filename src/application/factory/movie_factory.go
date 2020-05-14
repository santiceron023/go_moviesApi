package factory

import (
	"movies/src/application/command"
	"movies/src/domain/model"
)

func CreateMovie(command command.MovieCommand) model.Movie {
	var movie model.Movie
	movie, err := user.CreateUser(userCommand.FirstName, userCommand.LastName, userCommand.Email, userCommand.Password)
	return user, err
}
