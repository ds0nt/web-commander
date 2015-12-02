import React from 'react'
import MessageFactory from '../interface/messages-out'

export default class ChatInput extends React.Component {
  constructor(props) {
    super(props)
    this.messageFactory = new MessageFactory(this.props.room)
  }
  onSubmit(e) {
    e.preventDefault()
    let message = this.refs.chatInput.value
    this.refs.chatInput.value = ""
    this.messageFactory.send(message)
  }
  render() {
    return (
      <form onSubmit={(e) => this.onSubmit(e)}>
        <input className="ui chat input" type="text" ref="chatInput"/>
      </form>
    );
  }
}
