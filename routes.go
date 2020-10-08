package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func setupRoutes(enrutador *mux.Router) {

	enrutador.HandleFunc("/hola", func(w http.ResponseWriter, r *http.Request) {
		respuesta := "Hola mundo"
		_, err := getDB()
		if err == nil {
			json.NewEncoder(w).Encode(respuesta)
		} else {
			fmt.Println(err)
			json.NewEncoder(w).Encode("Error conectando" + err.Error())
		}
	}).Methods("GET")

}
