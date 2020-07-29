import React from 'react';
import configs from '../../configs/configs.json'
import strings from '../../configs/pt_BR'
import './ChatInput.css';

import Button from '../Templates/Button';
import Input from '../Templates/Input';

export default props => {
  return (
    <span className="chat-input">
      <Input keyup={props.keyup} placeholder={strings.placeholder.typeMessage} maxlength={configs.maxMsgLength}
        change={props.change} value={props.message} />
      <Button label={strings.button.send} click={props.send} />
      <Button label={strings.button.leave} off click={props.leave} />
    </span>
  )
}