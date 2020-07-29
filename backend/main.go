package main

import (
	"fmt"
	"intelchat/pkg/config"
	"intelchat/pkg/logger"
	"intelchat/pkg/websocket"
	"net/http"
)

// readConfigs read the configuration file
func readConfigs() {
	err := config.ReadConfiguration()
	if err != nil {
		panic(fmt.Sprintln("Error loading configs:", err))
	}
}

// setupLogs initiate log files
func setupLogs() {
	if config.Configuration.LogFiles {
		err := logger.Setup()
		if err != nil {
			panic(fmt.Sprintln("Error setting up logs:", err))
		}
	} else {
		logger.LogMsg("***Start***")
	}
}

// setupRoutes creates the pool for websockets and set the route for new connections
func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	websocket.SetupHandlers()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		websocket.ServeWs(pool, w, r)
	})

	// Serve frontend static built files
	if config.Configuration.ServeFiles {
		fs := http.FileServer(http.Dir(config.Configuration.StaticFilesPath))
		http.Handle("/", fs)
	}
}

func main() {
	readConfigs()
	setupLogs()

	setupRoutes()
	http.ListenAndServe(":"+config.Configuration.ServerPort, nil)
}
