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
	NotUnderstood string `json:"not_understood"`
	CliWelcome    string `json:"CLI_welcome"`
}

type Configuration struct {
	ApiKey         string
	LanguageString string `json:"language"`
	Language       *Language
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
	Conf = new(Configuration)
	Conf.Language = new(Language)

	readConfigFile()
	readApiKey()
	readLanguageFile(Conf.LanguageString)
}
