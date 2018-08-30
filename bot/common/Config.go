package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configFilePath = "conf/common.json"
const apikeyFilePath = "conf/api-key"
const languageFilePattern = "conf/language.%s.json"

var Conf *Configuration

type Language struct {
	NotUnderstood               string `json:"not_understood"`
	CliWelcome                  string `json:"CLI_welcome"`
	GenericError                string `json:"generic_error"`
	GenericAlive                string `json:"generic_alive"`
	UserWelcome                 string `json:"user_welcome"`
	IntroductionPreviousStory   string `json:"introduction_previous_story"`
	IntroductionNewStory        string `json:"introduction_new_story"`
	CommandShowStory            string `json:"command_show_story"`
	CommandShowStoryDescription string `json:"command_show_story_description"`
	CommandStartDescription     string `json:"command_start_description"`
	CommandHelp                 string `json:"command_help"`
	CommandHelpDescription      string `json:"command_help_description"`
	CommandHelpText             string `json:"command_help_text"`
	NextUsersTurn               string `json:"next_users_turn"`
	NotYourTurn                 string `json:"not_your_turn"`
	NotYetImplemented           string `json:"not_yet_implemented"`
	MessageTooShort             string `json:"message_too_short"`
	MessageTooLong              string `json:"message_too_long"`
	MessageMissingSpaces        string `json:"message_missing_spaces"`
	NotPermitted                string `json:"not_permitted"`
	FullStoryText               string `json:"full_story"`
}

type Configuration struct {
	ApiKey                       string
	LanguageString               string `json:"language"`
	Language                     *Language
	DataDirectory                string `json:"data_directory"`
	LogFile                      string
	StorageBackend               string `json:"storage_backend"`
	StorageBackendMemory         string
	StorageBackendLocal          string
	StorageBackendLocalStoryFile string
	FullStorySource              string `json:"full_story_source"`

	MessageMinLength int `json:"message_length_minimum"`
	MessageMaxLength int `json:"message_length_maximum"`
	MessageMinSpaces int `json:"message_spaces_minimum"`

	TgMessageMaxLength  int `json:"message_spaces_minimum"`
	TgMessagesPerSecond int `json:"telegram_messages_per_second_limit"`

	TgCommandStart       string
	TgCommandPrefix      string
	HelpMessageDelimeter string

	ErrorStorageBackendUnknown string
}

func newConfiguration() *Configuration {
	conf := new(Configuration)
	conf.Language = new(Language)

	conf.LogFile = "log.txt"

	conf.StorageBackendMemory = "memory"
	conf.StorageBackendLocal = "filesystem"
	conf.StorageBackendLocalStoryFile = "story.txt"

	conf.TgCommandStart = "start"
	conf.TgCommandPrefix = "/"

	conf.HelpMessageDelimeter = " - "

	conf.ErrorStorageBackendUnknown = "Unknown Storage Backend"

	return conf
}

func readConfigFile() {
	file, err := os.Open(configFilePath)
	Check(err)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(Conf)
	Check(err)

	file.Close()
}

func readApiKey() {
	dat, err := ioutil.ReadFile(apikeyFilePath)
	Check(err)
	Conf.ApiKey = string(dat)
}

func readLanguageFile(lang string) {
	filepath := fmt.Sprintf(languageFilePattern, lang)

	file, err := os.Open(filepath)
	Check(err)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(Conf.Language)
	Check(err)

	file.Close()
}

func ReadConfig() {
	Conf = newConfiguration()

	readConfigFile()
	readApiKey()
	readLanguageFile(Conf.LanguageString)
}
