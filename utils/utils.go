package utils

import (
	"net/url"
	"regexp"
	"sync"
)

var userFormat = make(map[int64]string)
var mu sync.RWMutex

func IsValidURL(link string) bool {
	_, err := url.ParseRequestURI(link)
	if err != nil {
		return false
	}

	re := regexp.MustCompile(`(youtube\.com|youtu\.be|instagram\.com|tiktok\.com|threads\.net|facebook\.com|twitter\.com|linkedin\.com|pinterest\.com)`)
	return re.MatchString(link)
}

func GetUserFormat(chatID int64) string {
	mu.RLock()
	defer mu.RUnlock()
	if format, ok := userFormat[chatID]; ok {
		return format
	}
	return "video"
}

func SetUserFormat(chatID int64, format string) {
	mu.Lock()
	defer mu.Unlock()
	userFormat[chatID] = format
}