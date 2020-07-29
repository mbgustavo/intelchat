import React from 'react'
import './ChatHistory.css'

import Message from './Message'

export default props => {
  // Array of messages
  const messages = props.chatHistory.map((msg, index) => {
    return <Message key={index} message={msg} />
  })

  return (
    <div className="chat-history">
      <div>
        {messages}
        <div ref={props.messagesEndRef} />
      </div>
    </div>
  )
}