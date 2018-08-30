package storage

import (
	. "../common"
	. "../logging"
	"io/ioutil"
	"os"
	"strings"
)

var story string
var contributors map[int64]string
var lastContributor int64

var storyFile *os.File

func getStoryMemory(update MessageUpdate) string {
	return story
}

func GetStory(update MessageUpdate) string {
	switch Conf.StorageBackend {
	case Conf.StorageBackendMemory:
		return getStoryMemory(update)
	case Conf.StorageBackendLocal:
		return getStoryMemory(update)
	default:
		Fatal(Conf.ErrorStorageBackendUnknown)
	}
	return Conf.Language.GenericError
}

func cleanMessage(update MessageUpdate, message string) (string, string) {
	message = strings.TrimSpace(message)
	storedMessage := message
	if GetStory(update) != "" {
		storedMessage = " " + message
	}
	return storedMessage, message
}

func distributeStory(message string) Messages {
	distributions := make(Messages, len(contributors))
	for contributor := range contributors {
		if contributor == lastContributor {
			continue
		}
		distributions = append(distributions, NewMessageFromId(contributor, message))
	}
	return distributions
}

func AddUserToStory(update MessageUpdate) bool {
	_, present := contributors[update.Message.Chat.ID]
	if !present {
		contributors[update.Message.Chat.ID] = update.Message.From.UserName
	}
	return !present
}

func appendStoryMemory(message string) {
	story += message
}

func appendStoryLocal(message string) {
	_, err := storyFile.WriteString(message)
	Check(err)
}

func AppendStory(update MessageUpdate, message string) Messages {
	storedMessage, message := cleanMessage(update, message)

	appendStoryMemory(storedMessage)
	switch Conf.StorageBackend {
	case Conf.StorageBackendMemory:
	case Conf.StorageBackendLocal:
		appendStoryLocal(storedMessage)
	default:
		Fatal(Conf.ErrorStorageBackendUnknown)
	}

	lastContributor = update.Message.Chat.ID

	AddUserToStory(update)

	return distributeStory(message)
}

func UserInTurn(userId int64) bool {
	return userId != lastContributor
}

func initializeMemoryStorage() {
	story = ""
	lastContributor = 0
	contributors = make(map[int64]string)
}

func initializeLocalStorage() {
	os.MkdirAll(Conf.DataDirectory, os.ModePerm)
	storyFilepath := Conf.DataDirectory + Conf.StorageBackendLocalStoryFile

	f, err := os.OpenFile(storyFilepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	dat, err := ioutil.ReadFile(storyFilepath)
	Check(err)
	story = string(dat)

	storyFile = f
}

func InitializeStorage() {
	initializeMemoryStorage()

	switch Conf.StorageBackend {
	case Conf.StorageBackendMemory:
	case Conf.StorageBackendLocal:
		initializeLocalStorage()
	default:
		Fatal(Conf.ErrorStorageBackendUnknown)
	}
}
