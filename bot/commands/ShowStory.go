package commands

import (
	. "../common"
	. "../storage"
	"fmt"
)

func ProcessShowStory(update MessageUpdate) Messages {
	story := GetShortStory(
		fmt.Sprintf(Conf.Language.FullStoryText, "%s", Conf.FullStorySource),
		false)
	if story == "" {
		story = Conf.Language.NoStoryYet
	}

	return NewMessages(update, story)
}
