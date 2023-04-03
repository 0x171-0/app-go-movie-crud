package main

import (
	// "fmt" // print to console
	"log" // for log
	"encoding/json" // encode data to json response
	"net/http" // create server
	"math/rand" // create movie id int
	"strconv" // provides functions for converting strings to basic data types and vice versa.
	"github.com/gorilla/mux" // provive powerful and flexible router for building HTTP services
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set header content type to json
	w.Header().Set("Content-Type", "application/json")
	// encode json
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	// Set header content type to json
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[:index]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func getMovieById(w http.ResponseWriter, r *http.Request){
	// set Header
	w.Header().Set("Content-Type", "appplication/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	// set Header
	w.Header().Set("Content-Type", "appplication/json")

	var movie Movie
	// get data from body
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// generate id
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	// append to db
	movies = append(movies, movie)

	// response
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	// set Header
	w.Header().Set("Content-Type", "appplication/json")
	// get param
	params := mux.Vars(r)	
	// loop over the movies, range
	// delete original one
	 for index, item := range(movies){
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			// get data from body
			_ = json.NewDecoder(r.Body).Decode(&movie)
			// generate id
			movie.ID = params["id"]
			// append to db
			movies = append(movies, movie)
			// response
			json.NewEncoder(w).Encode(movie)
		}
	 }
	
}

func main(){
	r := mux.NewRouter()

	r.HandleFunc("/", getMovies).Methods("GET")
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")

	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movies One",Director: &Director{Firstname:"John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "45455", Title: "Movies Two",Director: &Director{Firstname:"Steve", Lastname: "Smith"}})

	log.Printf("Starting server at port 8080\n")
	log.Fatal(http.ListenAndServe(":8080", r))
}