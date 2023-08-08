package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var movies []Movie

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "345667", Title: "Movie one", Director: &Director{Firstname: "John", Lastname: "Nolan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "323445", Title: "Movie two", Director: &Director{Firstname: "Juan", Lastname: "Cholo"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server to port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
