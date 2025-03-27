package web

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Jeanga7/load_bot.git/config"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type Download struct {
	URL    string `json:"url"`
	Format string `json:"format"`
	Date   string `json:"date"`
}

func StartWebServer() {
	router := gin.Default()

	router.GET("/downloads", func(c *gin.Context) {
		db, err := sql.Open("sqlite3", config.SQLiteDBPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur DB"})
			return
		}
		defer db.Close()

		rows, err := db.Query("SELECT url, format, date FROM history ORDER BY date DESC")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur de requête"})
			return
		}
		defer rows.Close()

		var downloads []Download
		for rows.Next() {
			var d Download
			rows.Scan(&d.URL, &d.Format, &d.Date)
			downloads = append(downloads, d)
		}
		c.JSON(http.StatusOK, downloads)
	})

	// Le serveur écoute sur le port 8080
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Erreur lancement serveur web:", err)
	}
}
