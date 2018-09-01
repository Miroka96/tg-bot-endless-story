package common

type Empty struct{}

var Nil = Empty{}

type Store interface {
	GetStory() string
	AppendStory(string)
	AddChat(int64) bool // returns true, if the id was newly added
	GetChats() map[int64]Empty
	AddUser(string) bool // returns true, if the user was newly added
	GetUsers() map[string]Empty
	GetLastChat() int64
	SetLastChat(int64)
}
