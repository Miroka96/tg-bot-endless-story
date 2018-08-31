package common

type Empty struct{}

var Nil = Empty{}

type Store interface {
	GetStory() string
	AppendStory(string)
	AddChat(int64) bool
	GetChats() map[int64]Empty
	AddUser(string) bool
	GetUsers() map[string]Empty
	GetLastChat() int64
	SetLastChat(int64)
}
