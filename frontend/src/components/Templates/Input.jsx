import React from 'react'
import './Input.css'

export default props => {
  return (
    <input className= "input" onKeyUp={props.keyup} placeholder={props.placeholder} maxLength={props.maxlength}
      onChange={props.change} value={props.value} />
  )
}