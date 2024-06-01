package util

import (
	"log"
	"os"
)
// THis logger writes to a file called api.log. It allows me to catch errors snd debug where 
// things cant be sent to the terminal.
func logger() *os.File {
	logFile, err := os.OpenFile("api.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	return logFile
}

// allows me to log any type of data, and as much as needed as long as they
// are separated by commas.
func WriteLog(message interface{}, arg ...interface{}) {
	logFile := logger()
	defer logFile.Close() // Close the log file when the application exits

	logger := log.New(logFile, "", log.LstdFlags)
	if len(arg) != 0 {
        logger.Println(message, arg)
	} else {
        logger.Println(message)
	}
	
}
