package main

import (
	. "./bot"
	. "./bot/common"
	"./bot/storage"
	"io"
	"log"
	"os"
)

func logIntoFile() {
	logfile, err := os.OpenFile(Conf.DataDirectory+Conf.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	Check(err)

	mw := io.MultiWriter(os.Stderr, logfile)
	log.SetOutput(mw)
}

func main() {
	println("Opening the Book...")

	ReadConfig()
	storage.InitializeStorage()
	logIntoFile()

	println(Conf.Language.CliWelcome)

	Run()
}
