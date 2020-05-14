package adapter

import (
	"github.com/gocql/gocql"
	"movies/src/domain/model"
	"movies/src/infrastructure/repository"
	"movies/src/infrastructure/utils/rest_errors"
)

const (
	queryGetMovie    = "SELECT id, user_id, title,length,synopsis,image_url FROM movies WHERE id=?;"
	queryInsertMovie = "INSERT into movies (id, user_id, title,length,synopsis,image_url) VALUES (?,?,?,?,?,?);"
	queryListMovie   = "SELECT id, user_id, title,length,synopsis,image_url FROM movies;"
)

func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (model.Movie, error)
	List() ([]model.Movie, error)
	Save(movie model.Movie) error
}

type dbRepository struct {
}

func (r *dbRepository) Save(m model.Movie) error {
	if err := repository.GetSession().Query(queryInsertMovie,
		gocql.TimeUUID(),
		m.User_id,
		m.Title,
		m.Length,
		m.Synopsis,
		m.Image_url,
	).Exec(); err != nil {
		//TODO ERRR
		return nil
	}
	return nil

}

func (r *dbRepository) List() ([]model.Movie, error) {
	results := make([]model.Movie, 0)
	var result model.Movie
	iter := repository.GetSession().Query(queryListMovie).Iter()
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

func (r *dbRepository) GetById(id string) (model.Movie, error) {

	//r.Save(model.Movie{Title: "pelicula envidad desde go"})

	////a,_ := r.List()
	////fmt.Println(a)
	//return model.Movie{},nil

	var result model.Movie
	if err := repository.GetSession().Query(queryGetMovie, id).Scan(
		&result.Id,
		&result.User_id,
		&result.Title,
		&result.Length,
		&result.Synopsis,
		&result.Image_url,
	); err != nil {
		if err == gocql.ErrNotFound {
			return model.Movie{}, rest_errors.NewNotFoundError("no access token found with given id")
		}
		//TODO TERMINAR ERRORES PERSONALIZADOS
		return model.Movie{}, rest_errors.NewInternalServerError("error when trying to get current id", err)
	}
	return result, nil
}
