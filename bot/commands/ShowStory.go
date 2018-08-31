package commands

import (
	. "../common"
	. "../storage"
)

func ProcessShowStory(update MessageUpdate) Messages {
	story := GetShortStory()
	if story == "" {
		story = Conf.Language.NoStoryYet
	}

	return NewMessages(update, story)
}
