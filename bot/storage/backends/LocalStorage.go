package backends

import (
	. "../../common"
	. "../common"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type LocalData struct {
	storyFile    *os.File
	chatFile     *os.File
	userFile     *os.File
	lastChatFile *os.File
}

func getFilePath(filename string) string {
	return Conf.DataDirectory + filename
}

func readDataFile(filename string) string {
	dat, err := ioutil.ReadFile(getFilePath(filename))
	Check(err)
	return string(dat)
}

func readLines(filename string) []string {
	return strings.Split(readDataFile(filename), "\n")
}

func openDataFile(filename string) *os.File {
	filePath := getFilePath(filename)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	Check(err)
	return file
}

func NewLocalData() *LocalData {
	os.MkdirAll(Conf.DataDirectory, os.ModePerm)

	data := LocalData{
		storyFile:    openDataFile(Conf.StorageBackendLocalStoryFile),
		chatFile:     openDataFile(Conf.StorageBackendLocalChatFile),
		userFile:     openDataFile(Conf.StorageBackendLocalUserFile),
		lastChatFile: openDataFile(Conf.StorageBackendLocalLastChatFile),
	}
	return &data
}

func (data *LocalData) GetStory() string {
	return string(readDataFile(Conf.StorageBackendLocalStoryFile))
}

func appendToFile(file *os.File, str string) {
	_, err := file.WriteString(str)
	Check(err)
	err = file.Sync()
	Check(err)
}

func overwriteFile(file *os.File, str string) {
	file.Truncate(0)
	file.Seek(0, 0)
	appendToFile(file, str)
}

func intToDec(i int64) string {
	return strconv.FormatInt(i, 10)
}

func decToInt(dec string) int64 {
	i, err := strconv.ParseInt(dec, 10, 64)
	Check(err)
	return i
}

func (data *LocalData) AppendStory(message string) {
	appendToFile(data.storyFile, message)
}

func (data *LocalData) AddChat(chatId int64) bool {
	idString := intToDec(chatId) + "\n"
	appendToFile(data.chatFile, idString)
	return true // assumes to chat was not seen before
}

func (data *LocalData) AddUser(username string) bool {
	str := username + "\n"
	appendToFile(data.userFile, str)
	return true // assumes to user was not seen before
}

func (data *LocalData) GetChats() map[int64]Empty {
	lines := readLines(Conf.StorageBackendLocalChatFile)
	chats := make(map[int64]Empty, len(lines))

	for _, line := range lines {
		str := strings.TrimSpace(line)
		if str == "" {
			continue
		}
		chat := decToInt(str)
		chats[chat] = Nil
	}

	return chats
}

func (data *LocalData) GetUsers() map[string]Empty {
	lines := readLines(Conf.StorageBackendLocalUserFile)
	users := make(map[string]Empty, len(lines))

	for _, line := range lines {
		user := strings.TrimSpace(line)
		if user == "" {
			continue
		}
		users[user] = Nil
	}

	return users
}

func (data *LocalData) GetLastChat() int64 {
	dec := readDataFile(Conf.StorageBackendLocalLastChatFile)
	dec = strings.TrimSpace(dec)

	if dec == "" {
		return 0
	}
	return decToInt(dec)
}

func (data *LocalData) SetLastChat(chatId int64) {
	overwriteFile(data.lastChatFile, intToDec(chatId))
}
