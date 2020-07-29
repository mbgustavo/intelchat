package websocket

import (
	"intelchat/pkg/logger"

	"github.com/gorilla/websocket"
)

// Read is the reading loop for the client, reading the message received and procesing it in the adequate handler
func (c *Client) Read() {
	var msg wsMessage
	// Close situations that are expected to happen and should not be treated as unexpected errors
	var expectedClose = []int{
		websocket.CloseNormalClosure,
		websocket.CloseGoingAway, // Happens when the user leaves the page
		websocket.CloseNoStatusReceived, // Happens when socket is simply closed at client side
	}

	for {
		// Get the JSON received from the websocket
		if err := c.Conn.ReadJSON(&msg); err != nil {
			if websocket.IsCloseError(err, expectedClose...) {
				logger.LogMsg("websocket.client.Read: Client disconnected: %s", c.Nickname)
			} else {
				logger.LogError("websocket.client.Read: Error reading message: %+v", err)
			}
			return
		}
		logger.LogDebug("websocket.client.Read: Message Received: %+v", msg)

		// Search the adequate handler to process the message
		if h, found := Handlers[msg.Event]; found {
			err := h(c, msg.Body)
			if err != nil {
				logger.LogError("websocket.client.Read: Error handling message: %+v", err)
				return
			}
			
			// When an access message handling succeeds, it means that the client joined the pool
			if (msg.Event == EventAccess) { 
				defer func() {
					c.Pool.unregister <- c.Nickname
				}()
			}
		}
	}
}
