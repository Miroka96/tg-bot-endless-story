//package tg_bot_endless_story
package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

var token = os.Getenv("TBOT_API_KEY")

var story string = ""

func check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func startBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	check(err)

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)
	return bot
}

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	check(err)

	return updates
}

func processMessage(update tgbotapi.Update) tgbotapi.MessageConfig {
	story += " " + update.Message.Text

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, story)

	return msg
}

func main() {
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
