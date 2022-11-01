package main

import (
	"log"
	"os"
)

type Application struct {
	logInfo  *log.Logger
	logError *log.Logger
}

func InfoLogger() *log.Logger {
	logger := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	return logger
}

func ErrorLogger() *log.Logger {
	logger := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)
	return logger
}
