package bot

import (
	. "./commands"
	. "./common"
	"./storage"
	"strings"
)

func processCommand(update MessageUpdate) Messages {
	msg := update.Message.Text[1:]
	words := strings.Split(msg, " ")

	if len(words) == 0 {
		return NewMessages(update, Conf.Language.GenericError)
	}

	switch words[0] {
	case Conf.TgCommandStart:
		return ProcessStart(update)
	case Conf.Language.CommandShowStory:
		return ProcessShowStory(update)
	case "hi":
		return NewMessages(update, Conf.Language.GenericAlive)
	default:
		return ProcessHelp(update)
	}

}

func processMessage(update MessageUpdate) Messages {
	LogUpdate(update)
	storage.AddChatUser(update)

	if strings.HasPrefix(update.Message.Text, "/") {
		return processCommand(update)
	} else {
		return ProcessPlaintext(update)
	}
}
