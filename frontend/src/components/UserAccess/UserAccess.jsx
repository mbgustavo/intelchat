import React, { Component } from 'react'
import { connect, errorMsgs, socket } from '../../shared/apiSocket'
import strings from '../../configs/pt_BR'
import './UserAccess.css'

import AccessForm from './AccessForm'

export default class UserAccess extends Component {
  constructor(props) {
    super(props)

    // Close socket if a connection is open
    if (socket) socket.close()

    // Bind methods
    this.inputChange = this.inputChange.bind(this)
    this.keyup = this.keyup.bind(this)
    this.connect = this.connect.bind(this)

    // Initial state
    this.state = { nickname: '', errorMessage: '' }
  }

  /** inputChange is the handler for input change */
  inputChange(event) {
    this.setState({ nickname: event.target.value })
  }

  /** keyup is the handler for OnKeyUp event on input */
  keyup(event) {
    if (event.keyCode === 13) this.connect()
  }

  /** connect connects to websocket server */
  connect() {
    if (this.state.nickname) {
      connect(this.state.nickname)
        .then(users => {
          this.props.history.push({
            pathname: '/chat-room',
            state: { users, nickname: this.state.nickname }
          })
        })
        .catch(errorMessage => {
          this.setErrorMessage(errorMessage)
        })
    } else {
      this.setErrorMessage(errorMsgs.emptyNick)
    }
  }

  /** setErrorMessage maps an error message for a displayable message */
  setErrorMessage(errorMessage) {
    if (errorMessage === errorMsgs.nickUsed) {
      this.setState({ errorMessage: strings.nicknameUsed })
    } else if (errorMessage === errorMsgs.emptyNick) {
      this.setState({ errorMessage: strings.error.emptyNickname })
    } else if (errorMessage === errorMsgs.nickTooLarge) {
      this.setState({ errorMessage: strings.error.nicknameTooLarge })
    } else if (errorMessage === errorMsgs.fullRoom) {
      this.setState({ errorMessage: strings.error.fullRoom })
    } else {
      this.setState({ errorMessage: strings.error.unexpected })
    }
  }

  render() {
    return (
      <div className="user-access">
        <h1>{strings.welcome}</h1>
        <AccessForm keyup={this.keyup} connect={this.connect} nickname={this.state.nickname}
          change={this.inputChange} />
        <p className="error-message">{this.state.errorMessage}</p>
      </div>
    );
  };

}