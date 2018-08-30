package storage

import (
	. "../common"
	"./local"
	"./memory"
	"strings"
)

type Empty struct{}

var Nil = Empty{}

type Store interface {
	GetStory() string
	AppendStory(string)
	AddChat(int64) bool
	GetChats() map[int64]Empty
	AddUser(string) bool
	GetUsers() map[string]Empty
	GetLastChat() int64
	SetLastChat(int64)
}

var stores = map[string]Store{
	Conf.StorageBackendMemory: memory.NewMemoryData(),
	Conf.StorageBackendLocal:  local.NewLocalData(),
}

var readingStore Store
var writingStores []Store

func GetStory() string {
	return readingStore.GetStory()
}

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
	store.AddChat(update.Message.Chat.ID)
	store.AddUser(update.Message.From.UserName)
	store.SetLastChat(update.Message.Chat.ID)
}

func AppendStory(update MessageUpdate, message string) Messages {
	storedMessage, message := cleanMessage(message)

	readingStore.AppendStory(storedMessage)
	updateStatus(readingStore, update)

	for _, store := range writingStores {
		store.AppendStory(storedMessage)
		updateStatus(store, update)
	}

	return distributeText(message)
}

func IsUserInTurn(chatId int64) bool {
	return chatId != readingStore.GetLastChat()
}

func copyData(from Store, to Store) Store {
	to.AppendStory(from.GetStory())
	to.SetLastChat(from.GetLastChat())
	for chatId := range from.GetChats() {
		to.AddChat(chatId)
	}
	for username := range from.GetUsers() {
		to.AddUser(username)
	}
	return to
}

func InitializeStorage() {
	readingStore = stores[Conf.StorageBackendMemory]

	if Conf.StorageBackend != Conf.StorageBackendMemory {
		store, present := stores[Conf.StorageBackend]
		if !present {
			Fatal(Conf.ErrorStorageBackendUnknown)
		}
		writingStores = append(writingStores, store)
	}

	if len(writingStores) > 0 {
		copyData(writingStores[0], readingStore)
	}
}
