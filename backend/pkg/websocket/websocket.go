// Package websocket handle with websocket connections for a realtime chat
package websocket

import (
	"fmt"
	"intelchat/pkg/config"
	"intelchat/pkg/logger"
	"net/http"

	"github.com/gorilla/websocket"
)

// upgrade upgrades a http connection to a websocket connection
func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  config.Configuration.WebSocket.ReadBufferSize,
		WriteBufferSize: config.Configuration.WebSocket.WriteBufferSize,
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true } // Allow all connections

	return upgrader.Upgrade(w, r, nil) // Error should be checked on caller
}

// ServeWs defines the websocket endpoint
func ServeWs(pool *Pool, w http.ResponseWriter, r *http.Request) {
	logger.LogDebug("websocket.ServeWs: WebSocket Endpoint Hit")

	conn, err := upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
		logger.LogError("websocket.ServeWs: Error upgrading websocket: %+v", err)
		return
	}
	defer conn.Close()

	// Create Client and start the reading loop
	c := &Client{
		Conn: conn,
		Pool: pool,
	}

	c.Read()
}
