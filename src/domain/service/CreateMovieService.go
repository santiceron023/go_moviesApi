package service

import (
	"movies/src/domain/model"
	"movies/src/domain/port"
)

type ServiceCreateMovie interface {
	execute(model.Movie) error
}
type serviceCreateMovie struct {
	repositoryMovies port.MoviesRepository
	repositoryUsers  port.UsersRepository
}

func (s *serviceCreateMovie) Execute(movie model.Movie) error {
	if err := s.repositoryUsers.CheckUser(movie.User_id);
		err != nil {
		//TODO
		return nil
	}
	if err := s.repositoryUsers.UpdateMovieCount(movie.User_id);
		err != nil {
		//TODO
		return nil
	}

	if err := s.repositoryMovies.Save(movie);
		err != nil {
		//TODO
		return nil
	}

	return nil


}
