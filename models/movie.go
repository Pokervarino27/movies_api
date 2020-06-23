package models

//Movie estructura basica de una pelicula
type Movie struct {
	MovieID  int    `json:"movieId"`
	Name     string `json:"name"`
	Year     int    `json:"year"`
	Director string `json:"director"`
}

//Movies arreglo de structs movies.
// type Movies []Movie

//Message estrcutura que retorna un mensaje de respuesta
// type Message struct {
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// }
