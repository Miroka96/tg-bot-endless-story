package common

import "github.com/go-telegram-bot-api/telegram-bot-api"

var story string = ""
var contributors = make(map[int64]string)
var lastContributor int64

func GetStory(update tgbotapi.Update) string {
	return story
}

func AppendStory(update tgbotapi.Update, message string) []tgbotapi.MessageConfig {
	if story != "" {
		story += " "
	}
	story += message
	lastContributor = update.Message.Chat.ID

	_, present := contributors[lastContributor]
	if !present {
		contributors[lastContributor] = update.Message.From.UserName
	}

	distributions := make([]tgbotapi.MessageConfig, 0, len(contributors))
	for contributor := range contributors {
		if contributor == lastContributor {
			continue
		}
		distributions = append(distributions, tgbotapi.NewMessage(contributor, message))
	}

	return distributions
}

func UserInTurn(userId int64) bool {
	return userId != lastContributor
}
