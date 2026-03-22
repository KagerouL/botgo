package main

import (
	"log"
	"os"
)

var messageLogger *log.Logger

func initLogger() {

	os.MkdirAll("logs", os.ModePerm)

	file, err := os.OpenFile(
		"logs/messages.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)

	if err != nil {
		panic(err)
	}

	messageLogger = log.New(file, "", log.Ldate|log.Ltime)
}