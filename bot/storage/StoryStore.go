package storage

import (
	. "../common"
	"./backends"
	. "./common"
)

var stores map[string]Store

var readingStore Store
var writingStores []Store

func keepChatUser(store Store, update MessageUpdate, keep ...bool) (bool, bool) {
	var newChat, newUser bool

	if len(keep) < 1 || keep[0] {
		newChat = store.AddChat(update.Message.Chat.ID)
	}
	if len(keep) < 2 || keep[1] {
		newUser = store.AddUser(update.Message.From.UserName)
	}

	return newChat, newUser
}

func AddChatUser(update MessageUpdate) {
	newChat, newUser := keepChatUser(readingStore, update)

	for _, store := range writingStores {
		keepChatUser(store, update, newChat, newUser)
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
