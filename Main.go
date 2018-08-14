package main

import (
	"tg-bot-endless-story/common"
)

func main() {
	println("Opening the Book...")

	common.ReadConfig()
	common.InitializeStorage()

	println(common.Conf.Language.CliWelcome)

	common.Run()
}
