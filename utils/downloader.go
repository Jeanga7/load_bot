package utils

import (
	"fmt"
	"os"
	"os/exec"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func DownloadAndSend(chatID int64, url string, format string, bot *tgbotapi.BotAPI) error {
	outputFile := "downloads/output.mp4"
	if format == "audio" {
		outputFile = "downloads/output.mp3"
	}

	// Prépare les arguments pour yt-dlp
	cmdArgs := []string{"-o", outputFile, "-f", "best"}
	if format == "audio" {
		cmdArgs = append(cmdArgs, "--extract-audio", "--audio-format", "mp3")
	}
	cmdArgs = append(cmdArgs, url)

	cmd := exec.Command("yt-dlp", cmdArgs...)

	// Lancer la commande
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erreur: %s\nsortie: %s", err.Error(), string(output))
	}

	// Envoi du fichier téléchargé
	if format == "audio" {
		audioMsg := tgbotapi.NewAudio(chatID, tgbotapi.FilePath(outputFile))
		_, err = bot.Send(audioMsg)
	} else {
		videoMsg := tgbotapi.NewVideo(chatID, tgbotapi.FilePath(outputFile))
		_, err = bot.Send(videoMsg)
	}
	// Suppression du fichier temporaire
	os.Remove(outputFile)
	if err != nil {
		return err
	}
	return nil
}

