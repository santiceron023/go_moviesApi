package controller

import (
	"github.com/gin-gonic/gin"
	"movies/src/application/command"
	"movies/src/application/handler"
	"movies/src/infrastructure/utils/rest_errors"
	"net/http"
)

type MoviesRestController interface {
	Get(*gin.Context)
	List(*gin.Context)
	Save(*gin.Context)
}

type movieRestController struct {
	getHandler    handler.GetHandler
	listHandler   handler.ListHandler
	createHandler handler.CreateHandler
}

func (rest *movieRestController) Save(context *gin.Context) {
	var movieCommand command.MovieCommand
	if errJson := context.ShouldBindJSON(&movieCommand); errJson != nil {
		restErr := rest_errors.NewBadRequestError("invalid request object")
		context.JSON(restErr.Status(),restErr)
		return
	}
	if errCreate := rest.createHandler.Execute(movieCommand); errCreate != nil {
		context.Error(errCreate)
		return
	}
	context.JSON(http.StatusOK,"")
}


func (rest *movieRestController) Get(context *gin.Context) {
	movieId := context.Param("movie_id")
	movie, err := rest.getHandler.Execute(movieId)
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, movie)
}

func (rest *movieRestController) List(context *gin.Context) {
	movies, err := rest.listHandler.Execute()
	if err != nil {
		context.Error(err)
		return
	}
	context.JSON(http.StatusOK, movies)
}

func NewRestHandler(getHandler handler.GetHandler, listHandler handler.ListHandler, createHandler handler.CreateHandler) *movieRestController {
	return &movieRestController{
		getHandler:    getHandler,
		listHandler:   listHandler,
		createHandler: createHandler,
	}
}
