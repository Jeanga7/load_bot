package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	BotToken      string
	SQLiteDBPath  string
	IsProduction  bool
)

func LoadConfig() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Avertissement : impossible de charger .env, on utilise les variables d'environnement")
	}

	BotToken = os.Getenv("BOT_TOKEN")
	
	SQLiteDBPath = os.Getenv("SQLITE_DB_PATH")
	if SQLiteDBPath == "" {
		SQLiteDBPath = "history.db"
	}

	if os.Getenv("ENV") == "production" {
		IsProduction = true
	} else {
		IsProduction = false
	}
}