package main

type Videojuego struct {
	Id     int64  `json:"id"`
	Nombre string `json:"nombre"`
	Genero string `json:"genero"`
	Anio   int64  `json:"anio"`
}
