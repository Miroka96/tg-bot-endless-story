package storage

import (
	. "../common"
	. "./common"
	"strings"
)

func cleanMessage(message string) (string, string) {
	message = strings.TrimSpace(message)
	storedMessage := message
	if readingStore.GetStory() != "" {
		storedMessage = " " + message
	}
	return storedMessage, message
}

func distributeText(messageText string) Messages {
	chats := readingStore.GetChats()
	lastChat := readingStore.GetLastChat()

	messages := make(Messages, len(chats))
	for chatId := range chats {
		if chatId == lastChat {
			continue
		}
		messages = append(messages, NewMessageFromId(chatId, messageText))
	}
	return messages
}

func updateStatus(store Store, update MessageUpdate) {
	store.SetLastChat(update.Message.Chat.ID)
}

func appendUpdateStory(store Store, update MessageUpdate, message string) {
	store.AppendStory(message)
	updateStatus(store, update)
}

func AppendStory(update MessageUpdate, message string) Messages {
	storedMessage, message := cleanMessage(message)

	appendUpdateStory(readingStore, update, storedMessage)

	for _, store := range writingStores {
		appendUpdateStory(store, update, storedMessage)
	}

	return distributeText(message)
}

func IsUserInTurn(chatId int64) bool {
	return chatId != readingStore.GetLastChat()
}
