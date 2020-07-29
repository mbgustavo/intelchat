package websocket

import (
	"errors"
	"intelchat/pkg/config"
	"intelchat/pkg/logger"
)

// ValidateNickname verify nickname conditions and returns an error if it can not be used
func ValidateNickname(c *Client) error {
	// Do not allow an empty nickname
	if c.Nickname == "" {
		logger.LogWarning("websocket.validators.ValidateNickname: Empty nickname received")
		return errors.New(ErrEmptyNick)
	}

	// Verify nickname max length
	if len(c.Nickname) > config.Configuration.MaxNickLength {
		logger.LogWarning("websocket.validators.ValidateNickname: Nickname above max length")
		return errors.New(ErrNickTooLarge)
	}

	// Verify if there is another user with the same nickname
	if _, used := c.Pool.Clients[c.Nickname]; used {
		logger.LogDebug("websocket.validators.ValidateNickname: Nickname already being used: %s", c.Nickname)
		return errors.New(ErrNickUsed)
	}

	return nil
}

// ValidatePoolConditions verify pool conditions for a client to join and returns an error if it can not
func ValidatePoolConditions(c *Client) error {
	// Verify if the chat room is full
	if len(c.Pool.Clients) >= config.Configuration.MaxClients {
		logger.LogDebug("websocket.validators.ValidatePoolConditions: Room full for user: %s", c.Nickname)
		return errors.New(ErrRoomFull)
	}

	return nil
}
