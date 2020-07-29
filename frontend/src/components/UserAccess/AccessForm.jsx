import React from 'react'
import configs from '../../configs/configs.json'
import strings from '../../configs/pt_BR'
import './AccessForm.css'

import Button from '../Templates/Button'
import Input from '../Templates/Input'

export default props => {
  return (
    <span className="access-form">
      <Input keyup={props.keyup} placeholder={strings.placeholder.chooseNickname} maxlength={configs.maxLength}
        change={props.change} value={props.nickname} />
      <Button label={strings.button.enter} click={props.connect} />
    </span>
  )
}