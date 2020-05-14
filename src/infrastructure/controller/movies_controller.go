package controller

import (
	"github.com/gin-gonic/gin"
	"movies/src/application/handler"
	"net/http"
)

type MoviesRestController interface {
	Get(c *gin.Context)
}

type movieRestController struct {
	getMovieHandler handler.GetHandler
}

func NewRestHandler(handlerser handler.GetHandler) *movieRestController {
	return &movieRestController{
		getMovieHandler: handlerser,
	}
}

func (m movieRestController) Get(c *gin.Context){
	movieId := c.Param("movie_id")
	movie, err := m.getMovieHandler.Handler(movieId)
	if err != nil{
		//TODO MEHORAR
		c.JSON(http.StatusNotImplemented,"la puta")
		return
	}
	c.JSON(http.StatusOK,movie)
}
