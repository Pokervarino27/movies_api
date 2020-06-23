package app

import (
	"movies_api/controllers"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

//StartApp funcion que inicializa aplicacion
func StartApp() {
	v1 := router.Group("/v1")
	{
		v1.GET("/movie/:id", controllers.GetMovie)
		v1.GET("/movies", controllers.ListMovies)
		v1.POST("/create", controllers.CreateMovie)
		v1.POST("/creates", controllers.CreateMovies)
		v1.POST("/rubies", controllers.CreateMoviesRuby)
	}

	if err := router.Run(":6767"); err != nil {
		panic(err)
	}
}
