package utils

import (
	"database/sql"

	"github.com/Jeanga7/load_bot.git/config"
	_ "github.com/mattn/go-sqlite3"
)

func CanDownload(chatID int64) bool {
	db, err := sql.Open("sqlite3", config.SQLiteDBPath)
	if err != nil {
		return false
	}
	defer db.Close()

	var count int
	// Compter le nombre de téléchargements du jour
	row := db.QueryRow("SELECT COUNT(*) FROM history WHERE chat_id = ? AND date >= date('now','start of day')", chatID)
	row.Scan(&count)
	// Par exemple : 5 téléchargements maximum par jour pour un utilisateur normal
	return count < 5
}
