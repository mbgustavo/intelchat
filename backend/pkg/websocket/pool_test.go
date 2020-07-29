package websocket_test

import (
	"intelchat/pkg/config"
	"intelchat/pkg/websocket"
	"strconv"
	"testing"
)

func TestGetNicknamesList(t *testing.T) {
	// Create pool with maximum of 10 clients
	config.Configuration.MaxClients = 10
	pool := websocket.NewPool()

	// Fulfill this pool with clients with nickname from 0 to 9
	for i := 0; i < config.Configuration.MaxClients; i++ {
		pool.Clients[strconv.Itoa(i)] = &websocket.Client{Nickname: strconv.Itoa(i), Pool: pool}
	}

	// Get all the nicknames
	nicknames := pool.GetNicknamesList()

	// Verify length of nicknames
	if len(nicknames) != config.Configuration.MaxClients {
		t.Errorf("list of nicknames sizes diverges: received %d, expected %d",
			len(nicknames), config.Configuration.MaxClients)
	}

	// All nicknames in the slice should be found in the clients map
	for _, nick := range nicknames {
		if _, ok := pool.Clients[nick]; !ok {
			t.Errorf("nickname not found: %s", nick)
		}
	}
}

func BenchmarkGetNicknamesList(b *testing.B) {
	// Create pool with maximum of 10 clients
	config.Configuration.MaxClients = 10
	pool := websocket.NewPool()

	// Fulfill this pool with clients with nickname from 0 to 9
	for i := 0; i < config.Configuration.MaxClients; i++ {
		pool.Clients[strconv.Itoa(i)] = &websocket.Client{Nickname: strconv.Itoa(i), Pool: pool}
	}

	for i := 0; i < b.N; i++ {
		pool.GetNicknamesList()
	}
}