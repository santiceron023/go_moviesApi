package port

import "movies/src/domain/model"

type MoviesRepository interface {
	GetById(movieId string) (model.Movie,error)
	List() ([]model.Movie,error)
	Save(movie model.Movie) error
}
