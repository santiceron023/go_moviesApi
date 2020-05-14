package handler

import (
	"movies/src/domain/model"
	"movies/src/domain/port"
)

type GetHandler interface {
	Execute(movieId string) (model.Movie,error)
}

type getHandler struct {
	movieRepository port.MoviesRepository
}

func NewGetHandler(db port.MoviesRepository) GetHandler {
	return &getHandler{
		movieRepository: db,
	}
}

func (handler *getHandler) Execute(movieId string)(model.Movie,error)  {
	movie, err := handler.movieRepository.GetById(movieId)
	if err != nil{
		return model.Movie{}, err
	}
	return movie,nil
}