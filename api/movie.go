package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	util "github.com/pokervarino27/movies_api/utils"
)

//Movie estructura basica de una pelicula
type Movie struct {
	MovieID  int    `json:"movieId"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
	Rating   string `json:"rating"`
}

//CreateMovie create a new movie's register and add in dataBase
func CreateMovie(c *gin.Context) {
	var movie Movie

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	redis := c.MustGet("redisConn").(util.Conn)
	data, err := json.Marshal(movie)
	if err != nil {
		fmt.Println(err)
	}
	key := fmt.Sprintf("movieById[%d]", movie.MovieID)
	result := redis.SetKey(key, data, 1000)
	if result == true {
		fmt.Println("Saved in Redis")
	} else {
		fmt.Println("Couldn't save in Redis")
		c.JSON(http.StatusBadRequest, gin.H{"Error": "couldn't save in Reds"})
	}
	c.JSON(http.StatusCreated, data)
}

//GetMovie get a movie from our database in Redis
func GetMovie(c *gin.Context) {

	movie := &Movie{}
	redis := c.MustGet("redisConn").(util.Conn)
	key := fmt.Sprintf("movieById[%s]", c.Param("id"))
	fmt.Println(key)
	result, _ := redis.GetKey(key, movie)
	if result == false {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Register not found"})
		return
	}
	c.JSON(http.StatusCreated, movie)
}
