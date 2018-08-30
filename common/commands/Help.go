package commands

import (
	. "../../common"
	. "../config"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

var helpMessage = ""

func ProcessHelp(update tgbotapi.Update) tgbotapi.MessageConfig {
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
