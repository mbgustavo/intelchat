import configs from '../configs/configs.json'

let socket

// events contain the event names expected
const events = {
  access : 'access',
  accessResult: 'access-result',
  message: 'message',
  exit: 'exit'
}

// error messages expected to receive from server
const errorMsgs = {
  emptyNick: 'empty-nickname',
  nickUsed: 'nickname-used',
  nickTooLarge: 'nickname-too-large',
  roomFull: 'room-full'
}

/** connect returns a promise with the result of the attempt to connect to websocket server */
function connect(username) {
  return new Promise((resolve, reject) => {
    socket = new WebSocket(`ws://${configs.serverAddress}:${configs.serverPort}/ws`)

    socket.onopen = () => {
      socket.send(JSON.stringify({ event: events.access, body: username }))
    }

    // Handles an access-result message only
    socket.onmessage = (msg) => {
      let data = JSON.parse(msg.data)

      if (data.event === events.accessResult) {
        if (data.body.result) {
          resolve(data.body.users)
        } else {
          reject(data.body.reason)
        }
      }
    }

    socket.onclose = (event) => {
      console.warn('Socket Closed Connection: ', event)
      reject(event)
    }

    socket.onerror = (error) => {
      console.error('Socket Error: ', error)
      reject(error)
    }
  })
}

export { socket, events, errorMsgs, connect };