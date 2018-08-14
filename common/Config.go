package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const configFilePath = "conf/config.json"
const apikeyFilePath = "conf/api-key"
const languageFilePattern = "conf/language.%s.json"

var Conf *Configuration

type Language struct {
	NotUnderstood                 string `json:"not_understood"`
	CliWelcome                    string `json:"CLI_welcome"`
	GenericError                  string `json:"generic_error"`
	GenericAlive                  string `json:"generic_alive"`
	UserWelcome                   string `json:"user_welcome"`
	IntroductionPreviousStory     string `json:"introduction_previous_story"`
	IntroductionNewStory          string `json:"introduction_new_story"`
	CommandShowStory              string `json:"command_show_story"`
	CommandShowStoryDescription   string `json:"command_show_story_description"`
	CommandShowSummary            string `json:"command_show_summary"`
	CommandShowSummaryDescription string `json:"command_show_summary_description"`
	NextUsersTurn                 string `json:"next_users_turn"`
	NotYourTurn                   string `json:"not_your_turn"`
	NotYetImplemented             string `json:"not_yet_implemented"`
}

type Configuration struct {
	ApiKey         string
	LanguageString string `json:"language"`
	Language       *Language
}

func newConfiguration() *Configuration {
	conf := new(Configuration)
	conf.Language = new(Language)
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
