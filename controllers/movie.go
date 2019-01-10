package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

type Controller struct{}

var movies []models.Movie

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetMovies(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie models.Movie
		movies = []models.Movie{}

		rows, err := db.Query("select * from movies")
		logFatal(err)

		defer rows.Close()

		for rows.Next() {
			rows.Scan(&movie.Id, &movie.Title, &movie.Genre, &movie.Year)
			logFatal(err)

			movies = append(movies, movie)
		}

		json.NewEncoder(w).Encode(movies)
	}
}
func (c Controller) GetMovie(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var movie models.Movie
		params := mux.Vars(r)

		rows := db.QueryRow("select * from movies where id=$1", params["id"])

		err := rows.Scan(&movie.Id, &movie.Title, &movie.Genre, &movie.Year)
		logFatal(err)

		json.NewEncoder(w).Encode(movie)

	}
}
func (c Controller) AddMovie(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var movie models.Movie
		var movieId int

		json.NewDecoder(r.Body).Decode(&movie)

		db.QueryRow("insert into movies (title, genre, year) values($1, $2, $3) RETURNING id;",
			movie.Title, movie.Genre, movie.Year).Scan(&movieId)

		json.NewEncoder(w).Encode(movieId)
	}
}

func (c Controller) UpdateMovie(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var movie models.Movie
		json.NewDecoder(r.Body).Decode(&movie)

		res, err := db.Exec("update movies set title=$1, genre=$2, year=$3 where id=$4 RETURNING id",
			&movie.Title, &movie.Genre, &movie.Year, &movie.Id)

		rowsUpdated, err := res.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveMovie(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		res, err := db.Exec("delete from movies where id = $1", params["id"])
		logFatal(err)

		rowsDeleted, err := res.RowsAffected()
		logFatal(err)

		json.NewEncoder(w).Encode(rowsDeleted)

	}
}
