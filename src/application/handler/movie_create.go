package handler

import (
	"movies/src/application/command"
	"movies/src/application/factory"
	"movies/src/domain/model"
	"movies/src/domain/port"
)

type CreateHandler interface {
	Execute(movie command.MovieCommand) (model.Movie, error)
}

type createHandler struct {
	movieRepository port.MoviesRepository
}

func (handler *createHandler) Execute(command command.MovieCommand) (model.Movie, error) {
	movie, err := factory.CreateMovie(command)
	if err != nil {
		return model.Movie{}, err
	}
	movie, err = handler.movieRepository.Save(movie)
	if err != nil {
		return model.Movie{}, err
	}
	return movie, nil
}

func NewCreateHandler(repository port.MoviesRepository) CreateHandler {
	return &createHandler{
		movieRepository: repository,
	}
}
