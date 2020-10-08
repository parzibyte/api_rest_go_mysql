package main

func insertarVideojuego(videojuego Videojuego) error {
	bd, err := obtenerBd()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO videojuegos (nombre, genero, anio) VALUES (?, ?, ?)", videojuego.Nombre, videojuego.Genero, videojuego.Anio)
	return err
}
