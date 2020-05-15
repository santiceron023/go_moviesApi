package app

import (
	"github.com/gin-gonic/gin"
	"movies/src/application/handler"
	"movies/src/infrastructure/controller"
	"movies/src/infrastructure/middleare"
	"movies/src/infrastructure/repository/adapter"
	"movies/src/infrastructure/repository/client"
)

var (
	router = gin.Default()
)

func StartApplication() {

	router.Use(middleare.ErrorHandler())

	movieRepository := adapter.NewRepository()
	getHandler := handler.NewGetHandler(movieRepository)
	listHandler := handler.NewListHandler(movieRepository)
	createHandler := handler.NewCreateHandler(movieRepository)
	restHandler := controller.NewRestHandler(getHandler,listHandler,createHandler)

	mapUrls(restHandler)
	client.GetSession()

	router.Run(":8080")
}
