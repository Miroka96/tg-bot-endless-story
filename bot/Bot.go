package bot

import (
	. "./common"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func startBot(token string) *tgbotapi.BotAPI {
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
	bot := startBot(Conf.ApiKey)
	updates := getUpdates(bot)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		responses := processMessage(MessageUpdate(update))

		for _, response := range responses {
			bot.Send(tgbotapi.MessageConfig(response))
		}
	}
}
