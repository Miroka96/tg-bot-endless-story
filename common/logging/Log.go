package logging

import "log"

func LogAuthorized(username string) {
	log.Printf("Authorized on account %s", username)
}
