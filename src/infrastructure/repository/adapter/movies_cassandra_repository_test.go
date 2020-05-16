package adapter

import (
	"context"
	"github.com/docker/go-connections/nat"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"log"
	"movies/src/domain/model"
	"movies/src/infrastructure/repository/client"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Starting cassandra container...")
	cassandraPort := nat.Port("9042/tcp")

	_, err := tc.GenericContainer(context.Background(),
		tc.GenericContainerRequest{
			ContainerRequest: tc.ContainerRequest{
				Image:        "cassandra",
				ExposedPorts: []string{cassandraPort.Port()},
				WaitingFor: wait.ForAll(
					wait.ForListeningPort(cassandraPort),
				),
			},
			Started: true, // auto-start the container
		})
	if err != nil {
		log.Fatal("starting container error:", err)
	}

	log.Println("container Ready")

	client.CreateKeyspace("meli")
	os.Exit(m.Run())
}
func TestCassandraRepository(t *testing.T) {
	client.CreateSession()
	client.CreateTable()

	repo := NewRepository()

	savedMovie,errSave := repo.Save(model.Movie{
		Id: "", User_id: 324, Title: "Die Hard", Length: 324, Synopsis: "a real action movie", Image_url: "",
	})
	assert.Nil(t, errSave)
	assert.NotEmpty(t,savedMovie.Id)

	listedMovie, errGet := repo.GetById(savedMovie.Id)
	assert.Nil(t, errGet)
	assert.NotEmpty(t,listedMovie)

	repo.Save(model.Movie{
		Id: "", User_id: 324, Title: "Pirate of the caribbean", Length: 324, Synopsis: "a real amazing adventure movie", Image_url: "",
	})
	movieList, errList := repo.List()
	assert.Nil(t, errList)
	assert.Equal(t,len(movieList),2)
}
