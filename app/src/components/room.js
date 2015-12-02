import React from 'react'
import Events from '../lib/event-master'

import ChatInput from './chat-input'
import ScriptBox from './script-box'

export default class Room extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      text: "Hello World",
      messages: [],
    }

    this.chatEvent = Events.roomOn(this.props.name, 'in:broadcast', (message) => this.addMessage(message))
    this.chatEvent2 = Events.roomOn(this.props.name, 'in:chat', (message) => this.addMessage(message))
    this.chatEvent3 = Events.roomOn(this.props.name, 'in:nick', (message) => this.addMessage(message))
    this.chatEvent4 = Events.roomOn(this.props.name, 'log', (message) => this.addMessage(message))
    this.chatEvent5 = Events.roomOn(this.props.name, 'socket:disconnect', () => this.addMessage("------------ You have been disconnected from the server -------------"))
  }

  componentWillUnmount() {
    this.chatEvent.off();
    this.chatEvent2.off();
    this.chatEvent3.off();
    this.chatEvent4.off();
    this.chatEvent5.off();
  }

  addMessage(message) {
    let messages = this.state.messages
    messages.push(message)
    this.setState({
      messages: messages
    })
  }

  renderMessage(message, k) {
    return (
      <div className="ui chat item" key={"message-" + k}>
        {message}
      </div>
    );
  }

  render() {
    let messages = this.state.messages.map(this.renderMessage)
    return (
      <div className="ui list messages main-pane" style={{display: this.props.active ? "block" : "none" }}>
        <div className="main-pane-inner">
          <div className="ui header item main-pane-item">{this.state.text}</div>
          {messages}
          <ChatInput room={this.props.name}/>
          <ScriptBox />
        </div>
      </div>
    );
  }
}
