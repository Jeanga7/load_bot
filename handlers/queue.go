package handlers

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DownloadTask struct {
	ChatID int64
	URL    string
	Format string
	Bot    *tgbotapi.BotAPI
}

var (
	taskQueue = make(chan DownloadTask, 20)
	wg        sync.WaitGroup
)

func StartQueueProcessor() {
	go func() {
		for task := range taskQueue {
			ProcessDownload(task.ChatID, task.URL, task.Format, task.Bot)
			wg.Done()
		}
	}()
}

func AddToQueue(chatID int64, url string, format string, bot *tgbotapi.BotAPI) {
	task := DownloadTask{
		ChatID: chatID,
		URL:    url,
		Format: format,
		Bot:    bot,
	}
	wg.Add(1)
	taskQueue <- task
}
