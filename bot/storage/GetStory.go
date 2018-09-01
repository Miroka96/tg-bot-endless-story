package storage

import (
	. "../common"
	"fmt"
	"strings"
)

func GetStory() string {
	return readingStore.GetStory()
}

func shortenStory(story string, maxLength int) string {
	storyWords := strings.SplitAfter(story, Conf.TextDelimeter)

	shortStory := ""
	length := 0

	for i := len(storyWords) - 1; i >= 0; i-- {
		wordLength := len([]rune(storyWords[i]))
		length += wordLength

		if length >= maxLength {
			startIndex := i
			if length > maxLength && wordLength < Conf.WordSplittingLength {
				startIndex++
			}
			shortStory = strings.Join(storyWords[startIndex:], "")

			if length > maxLength && wordLength >= Conf.WordSplittingLength {
				shortStory = shortStory[-maxLength:]
			}

			break
		}
	}
	return strings.TrimSpace(shortStory)
}

func wrapStoryUnderLimit(wrappingPattern string, story string) string {
	formatReplacementStringLength := 2 // %s
	maxLength := Conf.TgMessageMaxLength - len([]rune(wrappingPattern)) - formatReplacementStringLength
	storyLength := len([]rune(story))

	shortenedStoryPrefixLength := len([]rune(Conf.Language.ShortenedStoryPrefix))

	shortStory := Conf.Language.ShortenedStoryPrefix

	if storyLength <= maxLength {
		shortStory = story
	} else {
		shortStory += shortenStory(story, maxLength-shortenedStoryPrefixLength)
	}

	wrappedStory := fmt.Sprintf(wrappingPattern, shortStory)
	return wrappedStory
}

func GetShortStory(wrappingPattern string, forceWrapping bool) string {
	story := GetStory()
	if !forceWrapping && len([]rune(story)) <= Conf.TgMessageMaxLength {
		return story
	}

	return wrapStoryUnderLimit(wrappingPattern, story)
}
