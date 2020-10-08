package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	// Ping database
	bd, err := getDB()
	if err != nil {
		log.Printf("Error with database" + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error making connection to DB. Please check credentials. The error is: " + err.Error())
			return
		}
	}
	// Define routes
	router := mux.NewRouter()
	setupRoutesForVideoGames(router)
	// .. here you can define more routes
	// ...
	// fore xample setupRoutesForGenreS(router)

	// Setup and start server
	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr:    port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())
}
