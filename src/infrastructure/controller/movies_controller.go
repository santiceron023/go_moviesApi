package controller

import (
	"github.com/gin-gonic/gin"
	"movies/src/application/handler"
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
	//var movieCommand command.MovieCommand
	//if err := context.ShouldBindJSON(&movieCommand); err != nil {
	//	//todo ere
	//}
	//
	//if err := rest.createHandler.Execute(movieCommand);err != nil{
	//
	//}
	context.JSON(http.StatusOK,"ok")
}

func NewRestHandler(getHandler handler.GetHandler, listHandler handler.ListHandler, createHandler handler.CreateHandler) *movieRestController {
	return &movieRestController{
		getHandler:    getHandler,
		listHandler:   listHandler,
		createHandler: createHandler,
	}
}

func (rest *movieRestController) Get(c *gin.Context) {
	movieId := c.Param("movie_id")
	movie, err := rest.getHandler.Execute(movieId)
	if err != nil {
		//TODO MEHORAR
		c.JSON(http.StatusNotImplemented, "la puta")
		return
	}
	c.JSON(http.StatusOK, movie)
}

func (rest *movieRestController) List(c *gin.Context) {
	movie, err := rest.listHandler.Execute()
	if err != nil {
		//TODO MEHORAR
		c.JSON(http.StatusNotImplemented, "la puta")
		return
	}
	c.JSON(http.StatusOK, movie)
}
