package commands

import (
	. "../common"
	. "../storage"
)

func ProcessStart(update MessageUpdate) Messages {
	welcomeMessage := Conf.Language.StartPrefixUserWelcome
	story := GetStory()

	var response string

	if story == "" {
		response = welcomeMessage + Conf.Language.StartSuffixIntroductionNewStory
	} else {
		welcomeMessage += Conf.Language.StartSuffixIntroductionPreviousStory
		response = GetShortStory(welcomeMessage, true)
	}

	return append(NewMessages(update, response), ProcessHelp(update)...)
}
