package commands

import (
	. "../common"
	. "../storage"
)

func ProcessShowStory(update MessageUpdate) Messages {
	return NewMessages(update, GetStory(update))
}
