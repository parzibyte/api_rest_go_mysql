package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func obtenerBd() (*sql.DB, error) {
	return sql.Open("mysql", CadenaConexionBaseDeDatos)
}
