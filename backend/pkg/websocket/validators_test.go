package websocket_test

import (
	"intelchat/pkg/config"
	"intelchat/pkg/websocket"
	"strconv"
	"testing"
)

func TestValidateNickname(t *testing.T) {
	// Create client with empty nickname and try to validate, a warning will be logged here
	c := &websocket.Client{}
	if err := websocket.ValidateNickname(c); err == nil || err.Error() != websocket.ErrEmptyNick {
		t.Errorf("received %v, expected %v", err, websocket.ErrEmptyNick)
	}

	// Set max nickname length to 8 and try to validate with 9 characters, a warning will be logged here
	config.Configuration.MaxNickLength = 8
	c.Nickname = "123456789"
	if err := websocket.ValidateNickname(c); err == nil || err.Error() != websocket.ErrNickTooLarge {
		t.Errorf("received %v, expected %v", err, websocket.ErrNickTooLarge)
	}

	// Put a client with nickname 123 in the pool
	c.Pool = websocket.NewPool()
	c.Pool.Clients["123"] = &websocket.Client{Nickname: "123"}
	// Try to enter with the same nickname
	c.Nickname = "123"
	if err := websocket.ValidateNickname(c); err == nil || err.Error() != websocket.ErrNickUsed {
		t.Errorf("received %v, expected %v", err, websocket.ErrNickUsed)
	}

	// Change to a valid nickname and validation should succeed
	c.Nickname = "1234"
	if err := websocket.ValidateNickname(c); err != nil {
		t.Errorf("received %v, expected nil", err)
	}
}

func TestValidatePoolConditions(t *testing.T) {
	config.Configuration.MaxClients = 10
	pool := websocket.NewPool()
	c := &websocket.Client{Pool: pool}

	// Add 10 clients to the pool and try to join it
	for i := 0; i < config.Configuration.MaxClients; i++ {
		pool.Clients[strconv.Itoa(i)] = c
	}
	if err := websocket.ValidatePoolConditions(c); err == nil || err.Error() != websocket.ErrRoomFull {
		t.Errorf("received %v, expected %v", err, websocket.ErrRoomFull)
	}

	// Remove 1 client and the join should succeed
	delete(pool.Clients, "0")
	if err := websocket.ValidatePoolConditions(c); err != nil {
		t.Errorf("received %v, expected nil", err)
	}
}

func BenchmarkValidateNickname(b *testing.B) {
	pool := websocket.NewPool()
	c := &websocket.Client{Nickname: "123", Pool: pool} // create a client with a valid nickname
	for i := 0; i < b.N; i++ {
		websocket.ValidateNickname(c)
	}
}

func BenchmarkValidatePoolConditions(b *testing.B) {
	// create a pool with maximum of 10 clients and a client to join it
	config.Configuration.MaxClients = 10
	pool := websocket.NewPool()
	c := &websocket.Client{Pool: pool}

	for i := 0; i < b.N; i++ {
		websocket.ValidatePoolConditions(c)
	}
}
