package backends

import (
	. "../../common"
	. "../common"
	"io/ioutil"
	"os"
)

type LocalData struct {
	storyFilePath string
	storyFile     *os.File
}

func NewLocalData() LocalData {
	os.MkdirAll(Conf.DataDirectory, os.ModePerm)
	storyFilePath := Conf.DataDirectory + Conf.StorageBackendLocalStoryFile

	storyFile, err := os.OpenFile(storyFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)

	data := LocalData{
		storyFilePath: storyFilePath,
		storyFile:     storyFile,
	}
	return data
}

func (data LocalData) GetStory() string {
	dat, err := ioutil.ReadFile(data.storyFilePath)
	Check(err)
	return string(dat)
}

func (data LocalData) AppendStory(message string) {
	_, err := data.storyFile.WriteString(message)
	Check(err)
}

func (data LocalData) AddChat(chatId int64) bool {
	return false
}

func (data LocalData) AddUser(username string) bool {
	return false
}

func (data LocalData) GetChats() map[int64]Empty {
	return nil
}

func (data LocalData) GetUsers() map[string]Empty {
	return nil
}

func (data LocalData) GetLastChat() int64 {
	return 0
}

func (data LocalData) SetLastChat(chatId int64) {

}
