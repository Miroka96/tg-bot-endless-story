package common

import (
	. "./config"
	. "./logging"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var token string

func startBot() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(token)
	Check(err)

	bot.Debug = false

	LogAuthorized(bot.Self.UserName)
	return bot
}

func getUpdates(bot *tgbotapi.BotAPI) tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	Check(err)

	return updates
}

func Run() {
	token = Conf.ApiKey

	bot := startBot()
	updates := getUpdates(bot)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		responses := processMessage(update)

		for _, response := range responses {
			bot.Send(response)
		}
	}
}
