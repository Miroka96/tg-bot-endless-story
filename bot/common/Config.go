package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const configDirectory = "conf" + string(os.PathSeparator)
const configFilePath = configDirectory + "config.json"

var Conf *Configuration

type Language struct {
	NotUnderstood                        string `json:"not_understood"`
	CliWelcome                           string `json:"CLI_welcome"`
	GenericError                         string `json:"generic_error"`
	GenericAlive                         string `json:"generic_alive"`
	StartPrefixUserWelcome               string `json:"command_start_text_prefix_user_welcome"`
	StartSuffixIntroductionPreviousStory string `json:"command_start_text_previous_story"`
	StartSuffixIntroductionNewStory      string `json:"command_start_text_new_story"`
	CommandShowStory                     string `json:"command_show_story"`
	CommandShowStoryDescription          string `json:"command_show_story_description"`
	CommandStartDescription              string `json:"command_start_description"`
	CommandHelp                          string `json:"command_help"`
	CommandHelpDescription               string `json:"command_help_description"`
	CommandHelpText                      string `json:"command_help_text"`
	NextUsersTurn                        string `json:"next_users_turn"`
	NotYourTurn                          string `json:"not_your_turn"`
	NotYetImplemented                    string `json:"not_yet_implemented"`
	MessageTooShort                      string `json:"message_too_short"`
	MessageTooLong                       string `json:"message_too_long"`
	MessageMissingSpaces                 string `json:"message_missing_spaces"`
	NotPermitted                         string `json:"not_permitted"`
	FullStoryText                        string `json:"full_story"`
	NoStoryYet                           string `json:"no_story_yet"`
	ShortenedStoryPrefix                 string `json:"shortened_story_prefix"`
}

type Configuration struct {
	ApiKey                          string
	LanguageString                  string `json:"language"`
	Language                        *Language
	DataDirectory                   string `json:"data_directory"`
	LogFile                         string `json:"log_file"`
	StorageBackend                  string `json:"storage_backend"`
	StorageBackendMemory            string
	StorageBackendLocal             string
	StorageBackendLocalStoryFile    string `json:"storage_backend_local_story_file"`
	StorageBackendLocalChatFile     string `json:"storage_backend_local_chat_file"`
	StorageBackendLocalUserFile     string `json:"storage_backend_local_user_file"`
	StorageBackendLocalLastChatFile string `json:"storage_backend_local_lastchat_file"`
	FullStorySource                 string `json:"full_story_source"`
	LanguageFilePattern             string `json:"language_file_pattern"`
	ApiKeyFilePattern               string `json:"apikey_file_pattern"`
	Target                          string `json:"target"`

	MessageMinLength int `json:"message_length_minimum"`
	MessageMaxLength int `json:"message_length_maximum"`
	MessageMinSpaces int `json:"message_spaces_minimum"`

	TgMessageMaxLength  int `json:"telegram_message_length_maximum"`
	TgMessagesPerSecond int `json:"telegram_messages_per_second_limit"`

	TgCommandStart       string
	TgCommandPrefix      string
	HelpMessageDelimeter string `json:"help_message_delimeter"`
	WordDelimeter        string `json:"word_delimeter"`
	WordSplittingLength  int    `json:"word_splitting_length"`

	ErrorStorageBackendUnknown string
}

func newConfiguration() *Configuration {
	conf := new(Configuration)
	conf.Language = new(Language)

	conf.StorageBackendMemory = "memory"
	conf.StorageBackendLocal = "filesystem"

	conf.TgCommandStart = "start"
	conf.TgCommandPrefix = "/"

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
	apikeyFilePath := configDirectory + fmt.Sprintf(Conf.ApiKeyFilePattern, Conf.Target)
	dat, err := ioutil.ReadFile(apikeyFilePath)
	Check(err)
	apikey := string(dat)
	Conf.ApiKey = strings.TrimSpace(apikey)
}

func readLanguageFile(lang string) {
	filepath := configDirectory + fmt.Sprintf(Conf.LanguageFilePattern, lang)

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
