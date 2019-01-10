package main

import (
	"database/sql"
	"log"
	"net/http"

	"./controllers"
	"./driver"
	"./models"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
)

var movies []models.Movie
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()

	router := mux.NewRouter()

	controller := controllers.Controller{}

	router.HandleFunc("/movies", controller.GetMovies(db)).Methods("GET")
	router.HandleFunc("/movies/{id}", controller.GetMovie(db)).Methods("GET")
	router.HandleFunc("/movies", controller.AddMovie(db)).Methods("POST")
	router.HandleFunc("/movies", controller.UpdateMovie(db)).Methods("PUT")
	router.HandleFunc("/movies/{id}", controller.RemoveMovie(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
