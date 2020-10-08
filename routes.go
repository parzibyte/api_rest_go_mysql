package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutes(router *mux.Router) {

	router.HandleFunc("/videogames", func(w http.ResponseWriter, r *http.Request) {
		videoGames, err := getVideoGames()
		if err == nil {
			respondWithSuccess(videoGames, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods("GET")
	router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}
		videogame, err := getVideoGameById(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(videogame, w)
		}
	}).Methods("GET")

}

// Helper functions for respond with 200 or 500 code
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
