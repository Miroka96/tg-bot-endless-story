package main

import (
	"io"
	"log"
	"os"
	"tg-bot-endless-story/common"
)

func logIntoFile() {
	logfile, err := os.OpenFile(common.Conf.DataDirectory+common.Conf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	common.Check(err)

	mw := io.MultiWriter(os.Stderr, logfile)
	log.SetOutput(mw)
}

func main() {
	println("Opening the Book...")

	common.ReadConfig()
	common.InitializeStorage()
	logIntoFile()

	println(common.Conf.Language.CliWelcome)

	common.Run()
}
