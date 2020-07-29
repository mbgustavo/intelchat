import React from 'react'
import './User.css'

export default props => {
  let classes = 'user '
  classes += props.me ? 'me' : ''

  return (
    <div className={classes}>
      {props.nickname}
    </div>
  )
}