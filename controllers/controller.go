package controllers

import (
	"context"
	"fmt"
	"log"
	"movies_api/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getClient() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)

	err = client.Ping(context.TODO(), nil)
	fmt.Println("Connected to Mongo")
	return client
}

var collection = getClient().Database("testDb").Collection("movies")

//CreateMovie funcion que registra película en base de datos
func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Movie added: ", result.InsertedID)
	c.JSON(http.StatusCreated, result)
}

//CreateMovies crea multiples peliculas usando go routine
func CreateMovies(c *gin.Context) {
	var results []*models.Movie
	input := make(chan interface{})

	if err := c.ShouldBindJSON(&results); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	for _, current := range results {
		go CreateMoviesConcurrent(current, input)
		fmt.Println(current)
	}
	c.JSON(http.StatusCreated, results)
}

//CreateMoviesConcurrent maneja las go routine de guardado
func CreateMoviesConcurrent(movie *models.Movie, output chan interface{}) {
	insert, err := collection.InsertOne(context.TODO(), movie)
	if err != nil {
		output <- err
		log.Fatal(err)
		return
	}
	fmt.Println("Movie added: ", insert.InsertedID)
	output <- insert.InsertedID

}

//CreateMoviesRuby inserción bajo logica ruby
func CreateMoviesRuby(c *gin.Context) {
	var movies []*models.Movie
	if err := c.ShouldBindJSON(&movies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}

	for _, current := range movies {
		insert, err := collection.InsertOne(context.TODO(), current)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println("Movie added: ", insert.InsertedID)
	}
}

//GetMovie obtiene una película en base al ID
func GetMovie(c *gin.Context) {

	var movie models.Movie

	movieID, err := strconv.Atoi(c.Param("id"))
	fmt.Println(movieID)
	if err != nil {
		fmt.Printf("Error")
	}
	filter := bson.M{"movieid": movieID}
	err = collection.FindOne(context.TODO(), filter).Decode(&movie)
	if err != nil {
		// apiErr := utils.NotFoundAPIError("404 Not Found")
		c.JSON(http.StatusNotFound, 404)
		return
	}
	c.JSON(http.StatusOK, movie)
}

//ListMovies que retorna todas las películas en la API
func ListMovies(c *gin.Context) {

	var results []*models.Movie

	filter := bson.D{{}}
	options := options.Find()

	cur, err := collection.Find(context.TODO(), filter, options)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {
		var movie models.Movie
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &movie)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())
	c.JSON(http.StatusOK, results)

}
