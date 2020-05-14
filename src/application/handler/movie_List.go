package handler

import (
	"movies/src/domain/model"
	"movies/src/domain/port"
)

type ListHandler interface {
	Execute() ([]model.Movie,error)
}

type listHandler struct {
	movieRepository port.MoviesRepository
}

func NewListHandler(db port.MoviesRepository) ListHandler {
	return &listHandler{
		movieRepository: db,
	}
}

func (handler *listHandler) Execute()([]model.Movie,error)  {
	movie, err := handler.movieRepository.List()
	if err != nil{
		return make([]model.Movie,0), err
	}
	return movie,nil
}