/*

  ____          _____               _ _           _
 |  _ \        |  __ \             (_) |         | |
 | |_) |_   _  | |__) |_ _ _ __ _____| |__  _   _| |_ ___
 |  _ <| | | | |  ___/ _` | '__|_  / | '_ \| | | | __/ _ \
 | |_) | |_| | | |  | (_| | |   / /| | |_) | |_| | ||  __/
 |____/ \__, | |_|   \__,_|_|  /___|_|_.__/ \__, |\__\___|
         __/ |                               __/ |
        |___/                               |___/

____________________________________
/ Si necesitas ayuda, contáctame en \
\ https://parzibyte.me               /
 ------------------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
Creado por Parzibyte (https://parzibyte.me). Este encabezado debe mantenerse intacto,
excepto si este es un proyecto de un estudiante.
*/
package main

func createVideoGame(videoGame VideoGame) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", videoGame.Name, videoGame.Genre, videoGame.Year)
	return err
}

func deleteVideoGame(id int64) error {

	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM video_games WHERE id = ?", id)
	return err
}

// It takes the ID to make the update
func updateVideoGame(videoGame VideoGame) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE video_games SET name = ?, genre = ?, year = ? WHERE id = ?", videoGame.Name, videoGame.Genre, videoGame.Year, videoGame.Id)
	return err
}
func getVideoGames() ([]VideoGame, error) {
	//Declare an array because if there's error, we return it empty
	videoGames := []VideoGame{}
	bd, err := getDB()
	if err != nil {
		return videoGames, err
	}
	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT id, name, genre, year FROM video_games")
	if err != nil {
		return videoGames, err
	}
	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var videoGame VideoGame
		err = rows.Scan(&videoGame.Id, &videoGame.Name, &videoGame.Genre, &videoGame.Year)
		if err != nil {
			return videoGames, err
		}
		// and append it to the array
		videoGames = append(videoGames, videoGame)
	}
	return videoGames, nil
}

func getVideoGameById(id int64) (VideoGame, error) {
	var videoGame VideoGame
	bd, err := getDB()
	if err != nil {
		return videoGame, err
	}
	row := bd.QueryRow("SELECT id, name, genre, year FROM video_games WHERE id = ?", id)
	err = row.Scan(&videoGame.Id, &videoGame.Name, &videoGame.Genre, &videoGame.Year)
	if err != nil {
		return videoGame, err
	}
	// Success!
	return videoGame, nil
}
