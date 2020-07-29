import React from 'react'
import strings from '../../configs/pt_BR'
import './UsersList.css'

import User from './User'

export default props => {
  // Array of connected users
  const users = props.users.map((user) => {
    return <User key={user} me={props.nickname === user} nickname={user} />
  })

  return (
    <div className="users-list">
      <h3>{strings.users}</h3>
      <div>
        {users}
      </div>
    </div>
  )
}