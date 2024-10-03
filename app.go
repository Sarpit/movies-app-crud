package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialise INITIALISING FUNCTION
func (app *App) Initialise() error {

	connectionString := fmt.Sprintf("%v:%v@tcp(%v:3306)/%v", DBUser, DBPassword, DBHost, DBName)
	var err error
	app.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	app.Router = mux.NewRouter().StrictSlash(true)
	app.HandleRoutes()
	return nil
}

// Run Server
func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

// HandleRoutes Handling All Routes
func (app *App) HandleRoutes() {
	app.Router.HandleFunc("/", homepage).Methods("GET")
	app.Router.HandleFunc("/movies", app.getMovies).Methods("GET")
	app.Router.HandleFunc("/movies/{id}", app.getMovie).Methods("GET")
	app.Router.HandleFunc("/movies", app.createMovie).Methods("POST")
	app.Router.HandleFunc("/movies/{id}", app.updateMovie).Methods("PUT")
	app.Router.HandleFunc("/movies/{id}", app.deleteMovie).Methods("DELETE")
}

// sendResponse It sends the response
func sendResponse(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

// sendError It sends the error
func sendError(w http.ResponseWriter, statusCode int, err string) {
	errorMessage := map[string]string{"error": err}
	sendResponse(w, statusCode, errorMessage)
}

// DEFINING ALL ROUTES

// homepage Home page for movies CRUD application
func homepage(w http.ResponseWriter, r *http.Request) {
	log.Println("endpoint: homepage")
	fmt.Fprintf(w, "Welcomw to Movies CRUD Application")
}

// getMovies GET ALL MOVIES
func (app *App) getMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := getMovies(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, movies)

}

// getMovie GET MOVIE BY ID
func (app *App) getMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}

	var m movie
	m.ID = key

	err = m.getMovie(app.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			sendError(w, http.StatusInternalServerError, "No Movies exist with given ID")
			return
		default:
			sendError(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	sendResponse(w, http.StatusOK, m)

}

// createMovie CREATE MOVIE
func (app *App) createMovie(w http.ResponseWriter, r *http.Request) {
	var m movie
	err := json.NewDecoder(r.Body).Decode(&m)

	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid payload")
	}

	err = m.createMovie(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, m)

}

// updateMovie UPDATE MOVIE
func (app *App) updateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "invalid movie id")
	}

	var m movie

	err = json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}

	m.ID = key
	err = m.updateMovie(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
		return
	}

	sendResponse(w, http.StatusOK, m)
}

func (app *App) deleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		sendError(w, http.StatusBadRequest, "Invalid IP")
	}

	var m movie
	m.ID = key
	err = m.deleteMovie(app.DB)
	if err != nil {
		sendError(w, http.StatusInternalServerError, err.Error())
	}

	msg := map[string]string{"message": "Deleted movie"}
	sendResponse(w, http.StatusOK, msg)
}
