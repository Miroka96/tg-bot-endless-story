package common

import (
	. "./commands"
	. "./config"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

func processCommand(update tgbotapi.Update) []tgbotapi.MessageConfig {
	msg := update.Message.Text[1:]
	words := strings.Split(msg, " ")

	if len(words) == 0 {
		return []tgbotapi.MessageConfig{tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.GenericError)}
	}

	switch words[0] {
	case Conf.TgCommandStart:
		return ProcessStart(update)
	case Conf.Language.CommandShowStory:
		return []tgbotapi.MessageConfig{ProcessShowStory(update)}
	case "hi":
		return []tgbotapi.MessageConfig{tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.GenericAlive)}
	default:
		return []tgbotapi.MessageConfig{ProcessHelp(update)}
	}

}

func processMessage(update tgbotapi.Update) []tgbotapi.MessageConfig {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if strings.HasPrefix(update.Message.Text, "/") {
		return processCommand(update)
	} else {
		return ProcessPlaintext(update)
	}
}
