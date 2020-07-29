import React from 'react'
import strings from '../../configs/pt_BR'
import './Header.css'

export default () => {
  return (
    <header className="header">
      <h1>{strings.title}</h1>
    </header>
  )
}