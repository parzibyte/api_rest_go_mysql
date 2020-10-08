package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"
var (
	CadenaConexionBaseDeDatos = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("usuario"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("puerto"),
		os.Getenv("nombre"))
)
