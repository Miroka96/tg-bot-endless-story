package common

import "github.com/go-telegram-bot-api/telegram-bot-api"

type MessageUpdate tgbotapi.Update
type Message tgbotapi.MessageConfig
type Messages []Message

func NewMessageFromId(chatId int64, text string) Message {
	return Message(tgbotapi.NewMessage(chatId, text))
}

func NewMessage(update MessageUpdate, text string) Message {
	return NewMessageFromId(update.Message.Chat.ID, text)
}

func NewMessages(update MessageUpdate, messages ...string) Messages {
	configs := make([]Message, len(messages))
	for _, msg := range messages {
		configs = append(configs, NewMessage(update, msg))
	}
	return configs
}
