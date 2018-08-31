package backends

import (
	. "../common"
)

type MemoryData struct {
	story                string
	chatIds              map[int64]Empty
	contributors         map[string]Empty
	lastContributingChat int64
}

func NewMemoryData() MemoryData {
	return MemoryData{
		story:                "",
		chatIds:              make(map[int64]Empty),
		contributors:         make(map[string]Empty),
		lastContributingChat: 0,
	}
}

func (data MemoryData) GetStory() string {
	return data.story
}

func (data MemoryData) AppendStory(message string) {
	data.story += message
}

func (data MemoryData) AddChat(chatId int64) bool {
	_, present := data.chatIds[chatId]
	if !present {
		data.chatIds[chatId] = Nil
	}
	return !present
}

func (data MemoryData) AddUser(username string) bool {
	_, present := data.contributors[username]
	if !present {
		data.contributors[username] = Nil
	}
	return !present
}

func (data MemoryData) GetChats() map[int64]Empty {
	return data.chatIds
}

func (data MemoryData) GetUsers() map[string]Empty {
	return data.contributors
}

func (data MemoryData) GetLastChat() int64 {
	return data.lastContributingChat
}

func (data MemoryData) SetLastChat(chatId int64) {
	data.lastContributingChat = chatId
}
