package commands

import (
	. "../../common"
	. "../config"
	. "../storage"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func checkMessage(update tgbotapi.Update) []tgbotapi.MessageConfig {
	var resp string = ""

	msg := update.Message.Text
	spaces := strings.Count(msg, " ")
	l := len([]rune(msg))

	if l < Conf.MessageMinLength {
		resp = Conf.Language.MessageTooShort
	} else if l > Conf.MessageMaxLength {
		resp = Conf.Language.MessageTooLong
	} else if spaces < Conf.MessageMinSpaces {
		resp = Conf.Language.MessageMissingSpaces
	}

	if resp == "" {
		return nil
	} else {
		return []tgbotapi.MessageConfig{tgbotapi.NewMessage(update.Message.Chat.ID, resp)}
	}

}

func ProcessPlaintext(update tgbotapi.Update) []tgbotapi.MessageConfig {
	var responses []tgbotapi.MessageConfig

	responses = checkMessage(update)
	if responses != nil {
		return responses
	}

	if UserInTurn(update.Message.Chat.ID) {
		responses = AppendStory(update, update.Message.Text)
	} else {
		responses = []tgbotapi.MessageConfig{tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.NotYourTurn)}
	}

	return responses
}
