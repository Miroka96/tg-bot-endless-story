package commands

import (
	. "../../common"
	. "../storage"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func ProcessShowStory(update tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, GetStory(update))
}
