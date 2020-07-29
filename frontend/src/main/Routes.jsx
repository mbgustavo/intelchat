import React from 'react'
import { Switch, Route, Redirect } from 'react-router'

import UserAccess from '../components/UserAccess/UserAccess';
import ChatRoom from '../components/ChatRoom/ChatRoom';

export default () => {
  return (
    <Switch>
      <Route exact path="/" component={UserAccess} />
      <Route path="/chat-room" component={ChatRoom} />
      <Redirect from="*" to="/" />
    </Switch>
  )
}