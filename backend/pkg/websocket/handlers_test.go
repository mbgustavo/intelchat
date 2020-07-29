package websocket_test

import (
	"intelchat/pkg/websocket"
	"testing"
)

// Initial setup
func setup() *websocket.Client {
	websocket.SetupHandlers()
	pool := websocket.NewPool()
	go pool.Start()
	c := &websocket.Client{Pool: pool}
	return c
}

func TestMessageHandler(t *testing.T) {
	c := setup()

	// Get handler
	h, found := websocket.Handlers[websocket.EventMessage]
	if !found {
		t.Fatal("Handler not found: ", websocket.EventMessage)
	}

	// Execute it and don't expect any error
	err := h(c, "123")
	if err != nil {
		t.Fatal(err)
	}
}

func BenchmarkMessageHandler(b *testing.B) {
	c := setup()

	for i := 0; i < b.N; i++ {
		websocket.Handlers[websocket.EventMessage](c, "123")
	}
}