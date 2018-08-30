package common

import (
	"log"
	"strconv"
)

func LogAuthorized(username string) {
	log.Printf("Authorized on account %s", username)
}

func LogMessage(username string, chatId int64, msg string) {
	log.Printf("[%s %s] %s", username, strconv.FormatInt(chatId, 10), msg)
}

func LogUpdate(update MessageUpdate) {
	LogMessage(update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text)
}
