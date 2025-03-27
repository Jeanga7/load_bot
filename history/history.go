package history

import (
	"database/sql"
	"log"

	"github.com/Jeanga7/load_bot.git/config"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	db, err := sql.Open("sqlite3", config.SQLiteDBPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Créer la table si elle n'existe pas
	statement := `
	CREATE TABLE IF NOT EXISTS history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		chat_id INTEGER,
		url TEXT,
		format TEXT,
		date DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
}

func SaveDownload(chatID int64, url, format string) {
	db, err := sql.Open("sqlite3", config.SQLiteDBPath)
	if err != nil {
		log.Println("Erreur ouverture DB:", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO history (chat_id, url, format) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Erreur préparation statement:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(chatID, url, format)
	if err != nil {
		log.Println("Erreur insertion dans l'historique:", err)
	}
}
