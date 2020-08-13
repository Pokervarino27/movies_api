package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/pokervarino27/movies_api/api"
	"github.com/pokervarino27/movies_api/internal"
	utils "github.com/pokervarino27/movies_api/utils"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.New()
}

//StartApp funcion que inicializa aplicacion
func StartApp() {

	redis, _ := utils.RedisInit()

	router.Use(internal.RedisConnection(redis))
	v1 := router.Group("/v1")
	{
		v1.GET("/movie/:id", api.GetMovie)
		v1.POST("/create", api.CreateMovie)
	}

	if err := router.Run(":6767"); err != nil {
		panic(err)
	}
}
