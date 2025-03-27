package main

import (
	"log"

	"github.com/Jeanga7/load_bot.git/config"
	"github.com/Jeanga7/load_bot.git/handlers"
	"github.com/Jeanga7/load_bot.git/web"
)

func main() {
	config.LoadConfig()
	log.Println("🚀 Lancement du bot...")

	// demarrer le traitement de la file d'attente
	handlers.StartQueueProcessor()

	// Lance le serveur web en parallèle
	go web.StartWebServer()

	// Lance le bot Telegram
	handlers.StartTelegramBot(config.BotToken)
}