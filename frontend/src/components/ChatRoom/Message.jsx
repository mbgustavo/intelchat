import React from 'react'
import './Message.css'

export default props => {
  let message
  if (props.message.nickname) { // Message from a specific user
    message = <span><b>{props.message.nickname}:</b> {props.message.message}</span>
  } else { // Message with information
    message = <span className="info-message">{props.message.message}</span>
  }
  
  return (
    <div className="message">
      {props.message.time} {message}
    </div>
  )
}