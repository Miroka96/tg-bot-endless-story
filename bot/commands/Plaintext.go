package commands

import (
	. "../common"
	. "../storage"
	"strings"
)

func checkMessage(update MessageUpdate) Messages {
	var resp string = ""

	msg := update.Message.Text
	spaces := strings.Count(msg, " ")
	l := len([]rune(msg))

	if l < Conf.MessageMinLength {
		resp = Conf.Language.MessageTooShort
	} else if l > Conf.MessageMaxLength {
		resp = Conf.Language.MessageTooLong
	} else if spaces < Conf.MessageMinSpaces {
		resp = Conf.Language.MessageMissingSpaces
	}

	if resp == "" {
		return nil
	} else {
		return NewMessages(update, resp)
	}

}

func ProcessPlaintext(update MessageUpdate) Messages {
	var responses Messages

	responses = checkMessage(update)
	if responses != nil {
		return responses
	}

	if UserInTurn(update.Message.Chat.ID) {
		responses = AppendStory(update, update.Message.Text)
	} else {
		responses = NewMessages(update, Conf.Language.NotYourTurn)
	}

	return responses
}
