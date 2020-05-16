package adapter

import (
	"github.com/gocql/gocql"
	"movies/src/domain/exception"
	"movies/src/domain/model"
	"movies/src/infrastructure/repository/client"
	"movies/src/infrastructure/utils/rest_errors"
)

const (
	queryGetMovie    = "SELECT id, user_id, title,length,synopsis,image_url FROM movies WHERE id=?;"
	queryInsertMovie = "INSERT into movies (id, user_id, title,length,synopsis,image_url) VALUES (?,?,?,?,?,?);"
	queryListMovie   = "SELECT id, user_id, title,length,synopsis,image_url FROM movies;"
)

type CassandraRepository interface {
	GetById(string) (model.Movie, error)
	List() ([]model.Movie, error)
	Save(movie model.Movie) (model.Movie,error)
}

type cassandraRepository struct {
}

func NewRepository() CassandraRepository {
	return &cassandraRepository{}
}

func (r *cassandraRepository) Save(m model.Movie) (model.Movie,error) {
	id := gocql.TimeUUID()
	if err := client.GetSession().Query(queryInsertMovie,
		id,
		m.User_id,
		m.Title,
		m.Length,
		m.Synopsis,
		m.Image_url,
	).Exec(); err != nil {
		return model.Movie{},err
	}
	m.Id = id.String()
	return m,nil

}

func (r *cassandraRepository) List() ([]model.Movie, error) {
	results := make([]model.Movie, 0)
	var result model.Movie
	iter := client.GetSession().Query(queryListMovie).Iter()
	for iter.Scan(
		&result.Id,
		&result.User_id,
		&result.Title,
		&result.Length,
		&result.Synopsis,
		&result.Image_url,
	) {
		results = append(results, result)
		result = model.Movie{}
	}
	return results, nil
}

func (r *cassandraRepository) GetById(id string) (model.Movie, error) {

	var result model.Movie
	if err := client.GetSession().Query(queryGetMovie, id).Scan(
		&result.Id,
		&result.User_id,
		&result.Title,
		&result.Length,
		&result.Synopsis,
		&result.Image_url,
	); err != nil {
		if err == gocql.ErrNotFound {
			return model.Movie{}, exception.MovieNotFound{ErrMessage: "no movie found with given id"}
		}
		return model.Movie{}, rest_errors.NewInternalServerError("error when trying to get current id", err)
	}
	return result, nil
}



