package utils

import (
	"log"
)

func CheckError(err error, msg string) {
	if err != nil {
		log.Println(msg, err)
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
	  log.Fatalf("%s: %s", msg, err)
	}
}