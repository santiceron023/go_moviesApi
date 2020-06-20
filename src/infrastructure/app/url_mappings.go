package app

import "movies/src/infrastructure/controller"

func mapUrls(restcontrollers controller.MoviesRestController){
	router.GET("/movies/:movie_id",restcontrollers.Get)
	router.GET("/movies",restcontrollers.List)
	router.POST("/movies",restcontrollers.Save)
}
