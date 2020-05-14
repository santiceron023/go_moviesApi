package handler

import (
	"movies/src/domain/model"
	"movies/src/domain/port"
)

type CreateHandler interface {
	Execute(movie model.Movie) error
}

type createHandler struct {
	movieRepository port.MoviesRepository
}

func (handler *createHandler) Execute(movie model.Movie) error {
	if err := movie.Validate(); err != nil {
		return err
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
