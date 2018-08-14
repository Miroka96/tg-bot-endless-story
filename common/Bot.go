//package tg_bot_endless_story
package common

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

var token string //= os.Getenv("TBOT_API_KEY")

var story string = ""

func startBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	Check(err)

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	Check(err)

	return updates
}

func processCommand(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := update.Message.Text[1:]
	words := strings.Split(msg, " ")

	if len(words) == 0 {
		return tgbotapi.NewMessage(update.Message.Chat.ID, story)
	}

	return tgbotapi.NewMessage(update.Message.Chat.ID, "got it")
}

func processResponse(update tgbotapi.Update) tgbotapi.MessageConfig {
	story += " " + update.Message.Text

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, story)

	return msg
}

func processMessage(update tgbotapi.Update) tgbotapi.MessageConfig {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if strings.HasPrefix(update.Message.Text, "/") {
		return processCommand(update)
	} else {
		return processResponse(update)
	}
}

func Run() {
	token = Conf.ApiKey

	bot := startBot()
	updates := getUpdates(bot)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		response := processMessage(update)

		bot.Send(response)
	}
}
