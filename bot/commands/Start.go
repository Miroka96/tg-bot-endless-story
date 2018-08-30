package commands

import (
	. "../common"
	. "../storage"
	"fmt"
)

func ProcessStart(update MessageUpdate) Messages {
	var welcomeResponse = Conf.Language.UserWelcome
	welcomeResponse += "\n\n"
	story := GetStory()

	if story == "" {
		welcomeResponse += Conf.Language.IntroductionNewStory
	} else {
		welcomeResponse += fmt.Sprintf(Conf.Language.IntroductionPreviousStory, story)
	}

	return append(ProcessHelp(update), NewMessage(update, welcomeResponse))
}
