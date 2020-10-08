package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	// Primero hacerle un ping a la base de datos
	bd, err := obtenerBd()
	if err != nil {
		log.Printf("Error con la base de datos: " + err.Error())
		return
	} else {
		err = bd.Ping()
		if err != nil {
			log.Printf("Error conectando a la base de datos. Verifique los parámetros. El error es: " + err.Error())
			return
		}
	}
	insertarVideojuego(Videojuego{
		Id:     0,
		Nombre: "Mario xd",
		Genero: "No sé",
		Anio:   2012,
	})
	// Comenzamos con las rutas
	enrutador := mux.NewRouter()
	configurarRutas(enrutador)
	// Preparar y encender servidor
	direccion := ":8000"

	servidor := &http.Server{
		Handler: enrutador,
		Addr:    direccion,
		// Timeouts para evitar que el servidor se quede "colgado" por siempre
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Servidor iniciado y escuchando en el puerto %s", direccion)
	log.Fatal(servidor.ListenAndServe())
}
