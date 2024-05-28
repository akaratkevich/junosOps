package internal

import (
	"log"
	"os"
)

// Sets up logging to a specified file.
// The function takes a filename as a parameter and returns an error if it fails to open the file.

func SetupLogging(filename string) (*os.File, error) {
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil, err
	}
	log.SetOutput(logFile)
	return logFile, nil
}
