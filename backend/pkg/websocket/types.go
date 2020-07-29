package websocket

import (
	"github.com/gorilla/websocket"
)

// Event is an enum of strings for event names handled
type Event string

// Expected events
const (
	EventAccess       Event = "access"
	EventMessage            = "message"
	EventAccessResult       = "access-result"
	EventExit               = "exit"
)

// Error messages
const (
	ErrEmptyNick       = "empty-nickname"
	ErrNickTooLarge    = "nickname-too-large"
	ErrNickUsed        = "nickname-used"
	ErrRoomFull        = "room-full"
	ErrInvalidAsserion = "invalid-type-assertion"
)

// Client holds the user nickname, its websocket connection and the connection pool
type Client struct {
	Nickname string
	Conn     *websocket.Conn
	Pool     *Pool
}

// wsMessage is the struct used for websocket data send
type wsMessage struct {
	Event Event       `json:"event"`
	Body  interface{} `json:"body"`
}

// newUser is used for new access, with an error channel signaling its result
type newUser struct {
	client *Client
	err    chan error
}

// accessResult is the struct of the body sent to the client who tried to access the chat
type accessResult struct {
	Result bool     `json:"result"`
	Reason string   `json:"reason"`
	Users  []string `json:"users"` // slice with all nicknames of the logged users
}

// chatMessage is the struct of the body for a message in the chat
type chatMessage struct {
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
}
