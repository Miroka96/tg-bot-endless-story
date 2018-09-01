package storage

import (
	. "../common"
	"./backends"
	. "./common"
	"fmt"
	"strings"
)

var stores map[string]Store

var readingStore Store
var writingStores []Store

func GetStory() string {
	return readingStore.GetStory()
}

func shortenStory(story string, maxLength int) string {
	storyWords := strings.SplitAfter(story, Conf.TextDelimeter)

	shortStory := ""
	length := 0

	for i := len(storyWords) - 1; i >= 0; i-- {
		wordLength := len([]rune(storyWords[i]))
		length += wordLength

		if length >= maxLength {
			startIndex := i
			if length > maxLength && wordLength < Conf.WordSplittingLength {
				startIndex++
			}
			shortStory = strings.Join(storyWords[startIndex:], "")

			if length > maxLength && wordLength >= Conf.WordSplittingLength {
				shortStory = shortStory[-maxLength:]
			}

			break
		}
	}
	return strings.TrimSpace(shortStory)
}

func wrapStoryUnderLimit(wrappingPattern string, story string) string {
	formatReplacementStringLength := 2 // %s
	maxLength := Conf.TgMessageMaxLength - len([]rune(wrappingPattern)) - formatReplacementStringLength
	storyLength := len([]rune(story))

	shortenedStoryPrefixLength := len([]rune(Conf.Language.ShortenedStoryPrefix))

	shortStory := Conf.Language.ShortenedStoryPrefix

	if storyLength <= maxLength {
		shortStory = story
	} else {
		shortStory += shortenStory(story, maxLength-shortenedStoryPrefixLength)
	}

	wrappedStory := fmt.Sprintf(wrappingPattern, shortStory)
	return wrappedStory
}

func GetShortStory(wrappingPattern string, forceWrapping bool) string {
	story := GetStory()
	if !forceWrapping && len([]rune(story)) <= Conf.TgMessageMaxLength {
		return story
	}

	return wrapStoryUnderLimit(wrappingPattern, story)
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
