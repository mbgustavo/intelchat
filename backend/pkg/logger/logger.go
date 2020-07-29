// Package logger handle with event logging for the application
package logger

import (
	"intelchat/pkg/config"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

// logMux keeps the file writing thread-safe
var logMux sync.Mutex

// Setup configures the log files and create the directories if needed
func Setup() error {
	var err error

	// Create directory for the current day
	path := "log/" + time.Now().Format("2006-01-02") + "/"
	if _, err = os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
	}

	if err != nil {
		panic(err)
	}

	openLogFile(path + "main.log")

	return nil
}

// openLogFile returns a file for logging, with the append mode set, and also writes its first message
func openLogFile(filename string) *os.File {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	log.SetOutput(file)
	log.Print("***Start***")

	return file
}

// LogError prints an error message
func LogError(msg string, args ...interface{}) {
	prefix := "***Error***"
	LogMsg(strings.Join([]string{prefix, msg}, " "), args...)
}

// LogWarning prints a warning message
func LogWarning(msg string, args ...interface{}) {
	prefix := "***Warning***"
	LogMsg(strings.Join([]string{prefix, msg}, " "), args...)
}

// LogDebug prints a message if verbose logs mode is active
func LogDebug(msg string, args ...interface{}) {
	if config.Configuration.VerboseLogs {
		LogMsg(msg, args...)
	}
}

// LogMsg prints a message in the writer
func LogMsg(msg string, args ...interface{}) {
	logMux.Lock()
	defer logMux.Unlock()
	log.Printf(msg, args...)
}
