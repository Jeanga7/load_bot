package handlers

import (
	"fmt"
	"log"
	"time"

	"github.com/Jeanga7/load_bot.git/history"
	"github.com/Jeanga7/load_bot.git/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const maxRetries = 3

func ProcessDownload(chatID int64, url string, format string, bot *tgbotapi.BotAPI) {
	var err error
	for i := 0; i < maxRetries; i++ {
		bot.Send(tgbotapi.NewMessage(chatID, "📥 Téléchargement en cours ("+format+")... tentative "+fmt.Sprint(i+1)))
		err = utils.DownloadAndSend(chatID, url, format, bot)
		if err == nil {
			// Enregistrer dans l'historique
			history.SaveDownload(chatID, url, format)
			bot.Send(tgbotapi.NewMessage(chatID, "✅ Téléchargement terminé !"))
			return
		}
		log.Println("Erreur lors du téléchargement:", err)
		bot.Send(tgbotapi.NewMessage(chatID, "⚠ Tentative "+string(i+1)+" échouée, nouvelle tentative dans 5s..."))
		time.Sleep(5 * time.Second)
	}
	bot.Send(tgbotapi.NewMessage(chatID, "⚠ Échec du téléchargement après plusieurs tentatives."))
}
