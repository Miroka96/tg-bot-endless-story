package storage

import (
	. "../config"
	. "../logging"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"io/ioutil"
	"os"
	"strings"
)

var story string
var contributors map[int64]string
var lastContributor int64

var storyFile *os.File

func getStoryMemory(update tgbotapi.Update) string {
	return story
}

func GetStory(update tgbotapi.Update) string {
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

func cleanMessage(update tgbotapi.Update, message string) (string, string) {
	message = strings.TrimSpace(message)
	storedMessage := message
	if GetStory(update) != "" {
		storedMessage = " " + message
	}
	return storedMessage, message
}

func distributeStory(message string) []tgbotapi.MessageConfig {
	distributions := make([]tgbotapi.MessageConfig, 0, len(contributors))
	for contributor := range contributors {
		if contributor == lastContributor {
			continue
		}
		distributions = append(distributions, tgbotapi.NewMessage(contributor, message))
	}
	return distributions
}

func AddUserToStory(update tgbotapi.Update) bool {
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

func AppendStory(update tgbotapi.Update, message string) []tgbotapi.MessageConfig {
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
