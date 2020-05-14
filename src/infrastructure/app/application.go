package app

import (
	"github.com/gin-gonic/gin"
	"movies/src/application/handler"
	"movies/src/infrastructure/controller"
	"movies/src/infrastructure/repository"
	"movies/src/infrastructure/repository/adapter"
)

var (
	router = gin.Default()
)

func StartApplication() {

	//session, dbErr := cassandra.GetSesion()
	//if dbErr != nil {
	//	panic(dbErr)
	//}
	//session.Close()
	//atService := access_token.NewService(db.NewRepository())
	//atHandler := http.NewHandler(atService)
	movierepo := adapter.NewRepository()
	movieService := handler.NewGetHandler(movierepo)
	restHandler := controller.NewRestHandler(movieService)
	mapUrls(restHandler)
	repository.GetSession()
	router.Run(":8080")
}
