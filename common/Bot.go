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

func processStart(update tgbotapi.Update) []tgbotapi.MessageConfig {
	var welcomeResponse = Conf.Language.UserWelcome
	welcomeResponse += "\n\n"

	if GetStory(update) == "" {
		welcomeResponse += Conf.Language.IntroductionNewStory
	} else {
		welcomeResponse += fmt.Sprintf(Conf.Language.IntroductionPreviousStory, story)
	}

	return []tgbotapi.MessageConfig{
		tgbotapi.NewMessage(update.Message.Chat.ID, welcomeResponse),
		processHelp(update),
	}
}

func processShowStory(update tgbotapi.Update) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(update.Message.Chat.ID, GetStory(update))
}

var helpMessage = ""

func processHelp(update tgbotapi.Update) tgbotapi.MessageConfig {
	if helpMessage == "" {
		var helpActions strings.Builder
		partialStrings := []string{
			Conf.TgCommandStart, Conf.Language.CommandStartDescription,
			Conf.Language.CommandShowStory, Conf.Language.CommandShowStoryDescription,
			Conf.Language.CommandHelp, Conf.Language.CommandHelpDescription,
		}

		for i := 0; i < len(partialStrings); i++ {
			helpActions.WriteString(Conf.TgCommandPrefix)
			helpActions.WriteString(partialStrings[i])
			i++
			helpActions.WriteString(Conf.HelpMessageDelimeter)
			helpActions.WriteString(partialStrings[i])
			helpActions.WriteString("\n")
		}

		helpMessage = fmt.Sprintf(Conf.Language.CommandHelpText, helpActions.String())
	}

	return tgbotapi.NewMessage(update.Message.Chat.ID, helpMessage)
}

func processCommand(update tgbotapi.Update) []tgbotapi.MessageConfig {
	msg := update.Message.Text[1:]
	words := strings.Split(msg, " ")

	if len(words) == 0 {
		return []tgbotapi.MessageConfig{tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.GenericError)}
	}

	switch words[0] {
	case Conf.TgCommandStart:
		return processStart(update)
	case Conf.Language.CommandShowStory:
		return []tgbotapi.MessageConfig{processShowStory(update)}
	case "hi":
		return []tgbotapi.MessageConfig{tgbotapi.NewMessage(update.Message.Chat.ID, Conf.Language.GenericAlive)}
	default:
		return []tgbotapi.MessageConfig{processHelp(update)}
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

		responses := processMessage(update)

		for _, response := range responses {
			bot.Send(response)
		}
	}
}
