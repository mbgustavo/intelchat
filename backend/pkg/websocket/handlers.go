package websocket

import (
	"errors"
	"intelchat/pkg/logger"
)

// Handler is a function handler for a specific event, returning an error
type handler func(*Client, interface{}) error

// Handlers is a map where the event name is the key, and a handler function is the value
var Handlers map[Event]handler

// SetupHandlers creater the handlers map for all expected events
func SetupHandlers() {
	Handlers = map[Event]handler{
		EventMessage: handleMessage,
		EventAccess:  handleAccess,
	}
}

// handleMessage sends a message to all clients connected in the pool
func handleMessage(c *Client, data interface{}) error {
	// A string is expected in the body, assert that
	msgBody, ok := data.(string)
	if !ok {
		logger.LogError("websocket.handlers.handleMessage: Unexpected data: %+v", data)
		return errors.New(ErrInvalidAsserion)
	}

	// Update the body to a struct with message and nickname and broadcast it
	msg := wsMessage{Event: EventMessage, Body: chatMessage{Nickname: c.Nickname, Message: msgBody}}
	c.Pool.broadcast <- msg

	return nil
}

// handleAccess dels with a new client attempting to enter in the pool
func handleAccess(c *Client, data interface{}) error {
	// A string is expected in the body, assert that
	nickname, ok := data.(string)
	if !ok {
		logger.LogError("websocket.handlers.handleAccess: Unexpected data: %+v", data)
		return errors.New(ErrInvalidAsserion)
	}

	// Try to register the new user
	c.Nickname = nickname
	user := &newUser{client: c, err: make(chan error)}

	c.Pool.register <- user
	if err := <-user.err; err != nil {
		return err
	}

	logger.LogMsg("websocket.handlers.handleAccess: new user with nickname \"%s\"", nickname)
	return nil
}
