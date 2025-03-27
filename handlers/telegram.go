package handlers

import (
	"log"

	"github.com/Jeanga7/load_bot.git/utils"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func StartTelegramBot(token string) {
	var err error
	Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal("Erreur lors de la connexion au bot :", err)
	}

	Bot.Debug = true
	log.Println("🤖 Bot connecté :", Bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go handleMessage(update.Message)
		}
		// Gestion des callbacks (ex. pour le choix du format)
		if update.CallbackQuery != nil {
			go handleCallback(update.CallbackQuery)
		}
	}
}

func handleMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	text := msg.Text

	switch text {
	case "/start":
		response := "🎥 Bienvenue ! Envoie-moi un lien (ex. YouTube, Instagram, TikTok...) pour le télécharger."
		Bot.Send(tgbotapi.NewMessage(chatID, response))
		return
	case "/format":
		sendFormatChoice(chatID)
		return
	}

	// Vérifie que c’est une URL valide
	if !utils.IsValidURL(text) {
		Bot.Send(tgbotapi.NewMessage(chatID, "❌ Lien invalide. Envoie un lien supporté (ex. YouTube, Instagram, etc.)."))
		return
	}

	// Récupère le format choisi dans le contexte utilisateur (par défaut "video")
	format := utils.GetUserFormat(chatID)
	// Vérifie la limite d’usage (limiteur/VIP)
	if !utils.CanDownload(chatID) {
		Bot.Send(tgbotapi.NewMessage(chatID, "❌ Limite de téléchargements atteinte pour aujourd'hui."))
		return
	}

	Bot.Send(tgbotapi.NewMessage(chatID, "⏳ Ajout du lien à la file d'attente en format "+format+"..."))
	AddToQueue(chatID, text, format, Bot)
}

func sendFormatChoice(chatID int64) {
	buttons := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonData("📽 Vidéo", "format_video"),
		tgbotapi.NewInlineKeyboardButtonData("🎵 Audio", "format_audio"),
	}
	keyboard := tgbotapi.NewInlineKeyboardMarkup(buttons)
	msg := tgbotapi.NewMessage(chatID, "Choisissez le format souhaité :")
	msg.ReplyMarkup = keyboard
	Bot.Send(msg)
}

func handleCallback(callback *tgbotapi.CallbackQuery) {
	chatID := callback.Message.Chat.ID
	data := callback.Data

	if data == "format_video" || data == "format_audio" {
		format := "video"
		if data == "format_audio" {
			format = "audio"
		}
		utils.SetUserFormat(chatID, format)
		response := "✅ Format sélectionné : " + format + ". Maintenant, envoie-moi ton lien."
		edit := tgbotapi.NewEditMessageText(chatID, callback.Message.MessageID, response)
		Bot.Send(edit)
	}
}
