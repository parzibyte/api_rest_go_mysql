package main

func createVideoGame(videoGame VideoGame) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", videoGame.Name, videoGame.Genre, videoGame.Year)
	return err
}
