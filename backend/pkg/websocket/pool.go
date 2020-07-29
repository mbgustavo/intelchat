package websocket

import (
	"intelchat/pkg/config"
	"intelchat/pkg/logger"
)

// Pool is the struct with a map of all its clients and its necessary channels for communication
type Pool struct {
	Clients    map[string]*Client
	register   chan *newUser
	unregister chan string
	broadcast  chan wsMessage
}

// NewPool returns a new Pool pointer
func NewPool() *Pool {
	return &Pool{
		Clients:    make(map[string]*Client, config.Configuration.MaxClients),
		register:   make(chan *newUser),
		unregister: make(chan string),
		broadcast:  make(chan wsMessage),
	}
}

// GetNicknamesList get a list with all nicknames of the clients in the pool
func (pool *Pool) GetNicknamesList() []string {
	users := make([]string, len(pool.Clients))
	i := 0
	for nickname := range pool.Clients {
		users[i] = nickname
		i++
	}
	
	return users
}

// broadcastMessage sends a message to all clients in the pool
func (pool *Pool) broadcastMessage(e Event, body interface{}) {
	logger.LogDebug("websocket.pool.broadcastMessage: Sending message to all clients in Pool")
	for _, c := range pool.Clients {
		if err := c.Conn.WriteJSON(wsMessage{Event: e, Body: body}); err != nil {
			logger.LogError("websocket.pool.broadcastMessage: Message error: %+v", err)
		}
	}
}

// Start starts the pool listening for its channels, keeping the access of its clients thread-safe
func (pool *Pool) Start() {

	for {
		select {
		// New client attempting to join the pool
		case user := <-pool.register:
			if err := ValidatePoolConditions(user.client); err != nil {
				user.client.Conn.WriteJSON(wsMessage{
					Event: EventAccessResult,
					Body:  accessResult{Result: false, Reason: err.Error()},
				})
				user.err <- err
				continue
			}

			if err := ValidateNickname(user.client); err != nil {
				user.client.Conn.WriteJSON(wsMessage{
					Event: EventAccessResult,
					Body:  accessResult{Result: false, Reason: err.Error()},
				})
				user.err <- err
				continue
			}

			// This point down happens if the access was successful
			user.client.Conn.WriteJSON(wsMessage{
				Event: EventAccessResult,
				Body:  accessResult{Result: true, Users: pool.GetNicknamesList()},
			})

			user.err <- nil
			pool.Clients[user.client.Nickname] = user.client
			logger.LogMsg("websocket.pool: Size of Connection Pool: %d", len(pool.Clients))

			pool.broadcastMessage(EventAccess, chatMessage{Message: user.client.Nickname})

		// Client leaving the pool
		case nickname := <-pool.unregister:
			delete(pool.Clients, nickname)
			logger.LogMsg("websocket.pool: Size of Connection Pool: %d", len(pool.Clients))

			pool.broadcastMessage(EventExit, chatMessage{Message: nickname})

		// Message to broadcast
		case msg := <-pool.broadcast:
			pool.broadcastMessage(msg.Event, msg.Body)

		}
	}
}
