//package tg_bot_endless_story
package common

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strings"
)

var token string //= os.Getenv("TBOT_API_KEY")

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

func processStart(update tgbotapi.Update) tgbotapi.MessageConfig {
	var response = Conf.Language.UserWelcome
	response += "\n\n"

	if GetStory(update) == "" {
		response += Conf.Language.IntroductionNewStory
	} else {
		response += fmt.Sprintf(Conf.Language.IntroductionPreviousStory, story)
	}

	return tgbotapi.NewMessage(update.Message.Chat.ID, response)
}

func processShowStory(update tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, GetStory(update))
}

func processHelp(update tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.NotYetImplemented)
}

func processCommand(update tgbotapi.Update) tgbotapi.MessageConfig {
	msg := update.Message.Text[1:]
	words := strings.Split(msg, " ")

	if len(words) == 0 {
		return tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.GenericError)
	}

	switch words[0] {
	case "start":
		return processStart(update)
	case Conf.Language.CommandShowStory:
		return processShowStory(update)
	case "hi":
		return tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.GenericAlive)
	default:
		return processHelp(update)
	}

}

func processResponse(update tgbotapi.Update) []tgbotapi.MessageConfig {
	var responses []tgbotapi.MessageConfig
	if UserInTurn(update.Message.Chat.ID) {
		responses = AppendStory(update, update.Message.Text)
	} else {
		responses = []tgbotapi.MessageConfig{tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.NotYourTurn)}
	}

	return responses
}

func processMessage(update tgbotapi.Update) []tgbotapi.MessageConfig {
	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if strings.HasPrefix(update.Message.Text, "/") {
		return []tgbotapi.MessageConfig{processCommand(update)}
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

		responses := processMessage(update)

		for _, response := range responses {
			bot.Send(response)
		}
	}
}
