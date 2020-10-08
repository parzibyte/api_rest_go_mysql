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
	createVideoGame(VideoGame{
		Id:    0,
		Name:  "Mario xd",
		Genre: "No sé",
		Year:  2012,
	})
	// Define routes
	router := mux.NewRouter()
	setupRoutes(router)
	// Setup and start server
	direccion := ":8000"

	servidor := &http.Server{
		Handler: router,
		Addr:    direccion,
		// Timeouts para evitar que el servidor se quede "colgado" por siempre
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started at %s", direccion)
	log.Fatal(servidor.ListenAndServe())
}
