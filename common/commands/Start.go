package commands

import (
	. "../../common"
	. "../config"
	. "../storage"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func ProcessStart(update tgbotapi.Update) []tgbotapi.MessageConfig {
	var welcomeResponse = Conf.Language.UserWelcome
	welcomeResponse += "\n\n"
	story := GetStory(update)

	if story == "" {
		welcomeResponse += Conf.Language.IntroductionNewStory
	} else {
		welcomeResponse += fmt.Sprintf(Conf.Language.IntroductionPreviousStory, story)
	}

	return []tgbotapi.MessageConfig{
		tgbotapi.NewMessage(update.Message.Chat.ID, welcomeResponse),
		ProcessHelp(update),
	}
}
