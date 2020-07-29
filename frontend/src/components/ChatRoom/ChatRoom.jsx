import React, { Component } from 'react'
import { Prompt } from 'react-router'
import { socket, events } from '../../shared/apiSocket'
import strings from '../../configs/pt_BR'
import './ChatRoom.css'

import ChatHistory from './ChatHistory'
import UsersList from './UsersList'
import ChatInput from './ChatInput'

export default class ChatRoom extends Component {
  constructor(props) {
    super(props)

    // Prevent users from entering the room without appropriate connection
    if (!socket) this.leave()

    // Reference for the end of messages, used for auto-scroll
    this.messagesEndRef = React.createRef()

    // Bind methods
    this.inputChange = this.inputChange.bind(this)
    this.keyup = this.keyup.bind(this)
    this.send = this.send.bind(this)
    this.leave = this.leave.bind(this)

    // Handlers for websocket events
    this.handlers = {}
    this.handlers[events.message] = this.handleMessage.bind(this)
    this.handlers[events.access] = this.handleAccess.bind(this)
    this.handlers[events.exit] = this.handleExit.bind(this)

    // Initial state
    this.state = {
      chatHistory: [],
      users: [],
      message: '',
      nickname: '',
      confirmation: strings.confirmLeaveRoom
    }
  }

  componentDidMount() {
    if (socket) {
      if (this.props.location.state) {
        this.setState({ users: this.props.location.state.users, nickname: this.props.location.state.nickname })
      } else {
        console.error('State was not received')
      }

      // Set handlers for the event received
      socket.onmessage = (msg) => {
        let data = JSON.parse(msg.data)
        if (this.handlers[data.event]) {
          this.handlers[data.event](data.body)
        } else {
          console.warn('Event unknown: ', data.event)
        }
      }

      // If the connect is lost, back to access page
      socket.onclose = () => {
        this.setState({ confirmation: strings.connectionLost })
        this.leave()
      }

      socket.onerror = err => console.error('Socket Error: ', err)
    }
  }

  /** zeroFill standardizes numbers with leading zeroes. Ex.: 1 -> 01 */
  zeroFill(number) {
    if (number < 10) {
      return '0' + number
    }
    return number
  }

  /** getFormattedTime get the current time in format hh:mm:ss */
  getFormattedTime(now) {
    let hour = this.zeroFill(now.getHours())
    let minute = this.zeroFill(now.getMinutes())
    let second = this.zeroFill(now.getSeconds())
    return [hour, minute, second].join(':')
  }

  /** handleMessage handles with a new message in the chat */
  handleMessage(data) {
    data.time = this.getFormattedTime(new Date())

    this.setState(prevState => ({
      chatHistory: [...prevState.chatHistory, data]
    }))

    this.messagesEndRef.scrollIntoView({ behavior: 'smooth' })
  }

  /** handleAccess handles with the exit of an user from the chat */
  handleAccess(data) {
    let nickname = data.message
    let time = this.getFormattedTime(new Date())

    if (!this.state.users.includes(nickname)) {
      this.setState(prevState => ({
        chatHistory: [...prevState.chatHistory, { message: `${nickname} ${strings.newUser}`, time }],
        users: [...prevState.users, nickname]
      }))
    }

    this.messagesEndRef.scrollIntoView({ behavior: 'smooth' })
  }

  /** handleExit handles with the entrance of an user on the chat */
  handleExit(data) {
    let nickname = data.message
    let time = this.getFormattedTime(new Date())

    this.setState(prevState => ({
      chatHistory: [...prevState.chatHistory, { message: `${nickname} ${strings.userLeft}`, time }],
      users: prevState.users.filter(user => user !== nickname)
    }))

    this.messagesEndRef.scrollIntoView({ behavior: 'smooth' })
  }

  /** keyup is the handler for OnKeyUp event on input, sends when enter is pressed */
  keyup(event) {
    if (event.keyCode === 13) this.send()
  }

  /** inputChange is the handler for input change */
  inputChange(event) {
    this.setState({ message: event.target.value })
  }

  /** send sends message in the chat room */
  send() {
    if (this.state.message !== '') {
      socket.send(JSON.stringify({ event: events.message, body: this.state.message }))
      this.setState({ message: '' })
    }
  }

  /** leave returns to access page */
  leave() {
    this.props.history.push('/')
  }

  render() {
    return (
      <div className="chat-room">
        <Prompt message={this.state.confirmation} />
        <ChatHistory chatHistory={this.state.chatHistory} messagesEndRef={(el) => { this.messagesEndRef = el }} />
        <UsersList users={this.state.users} nickname={this.state.nickname} />
        <ChatInput change={this.inputChange} keyup={this.keyup} send={this.send} message={this.state.message}
          leave={this.leave} />
      </div>
    )
  }
}