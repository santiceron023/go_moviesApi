package handler

import (
	"movies/src/application/command"
	"movies/src/application/factory"
	"movies/src/domain/port"
)

type CreateHandler interface {
	Execute(movie command.MovieCommand) error
}

type createHandler struct {
	movieRepository port.MoviesRepository
}

func (handler *createHandler) Execute(command command.MovieCommand) error {
	movie, err := factory.CreateMovie(command)
	if err != nil {
		//todo error
	}
	if err := handler.movieRepository.Save(movie); err != nil {
		return err
	}
	return nil
}

func NewCreateHandler(repository port.MoviesRepository) CreateHandler {
	return &createHandler{
		movieRepository: repository,
	}
}
