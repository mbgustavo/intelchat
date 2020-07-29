import React from 'react'
import './Button.css'

export default props => {
  let classes = 'button '
  classes += props.off ? 'off' : ''

  return (
    <button className={classes} onClick={props.click}>{props.label}</button>
  )
}