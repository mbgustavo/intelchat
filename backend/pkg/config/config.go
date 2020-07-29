// Package config reads the configuration for the application from a JSON file
package config

import (
	"encoding/json"
	"os"
)

const filename = "./configs/configs.json"

// Configuration is structured to read the JSON file
var Configuration struct {
	ServerPort string `json:"serverPort"`
	WebSocket  struct {
		ReadBufferSize  int `json:"readBufferSize"`
		WriteBufferSize int `json:"writeBufferSize"`
	} `json:"webSocket"`
	LogFiles        bool   `json:"logFiles"`
	VerboseLogs     bool   `json:"verboseLogs"`
	MaxClients      int    `json:"maxClients"`
	MaxNickLength   int    `json:"maxNickLength"`
	ServeFiles			bool	 `json:"serveFiles"`
	StaticFilesPath string `json:"staticFilesPath"`
}

// ReadConfiguration read the JSON file and put the configuration in the Configuration struct
func ReadConfiguration() error {
	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Decode configurations
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&Configuration)
	if err != nil {
		return err
	}

	return nil
}
