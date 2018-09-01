package storage

import (
	. "../common"
	"./backends"
	. "./common"
)

var stores map[string]Store

var readingStore Store
var writingStores []Store

func keepChatUser(store Store, update MessageUpdate) {
	store.AddChat(update.Message.Chat.ID)
	store.AddUser(update.Message.From.UserName)
}

func AddChatUser(update MessageUpdate) {
	keepChatUser(readingStore, update)

	for _, store := range writingStores {
		keepChatUser(store, update)
	}
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

func initializeStores() {
	stores = map[string]Store{
		Conf.StorageBackendMemory: backends.NewMemoryData(),
		Conf.StorageBackendLocal:  backends.NewLocalData(),
	}
}

func InitializeStorage() {
	initializeStores()
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
