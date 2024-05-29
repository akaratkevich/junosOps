package internal

import (
	"log"
	"os"
)

func SetupLogging(filename string) (*os.File, error) {
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(logFile)
	return logFile, nil
}
