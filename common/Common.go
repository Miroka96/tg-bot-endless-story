package common

import "log"

func Check(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func Fatal(msg string) {
	log.Panic(msg)
}
