import Layout from './layout'
import eventbus from './event-master'
import { Message, MessageHandler } from './socket-master'

import ReactDOM from 'react-dom'
import React from 'react'


let addMessage = () => {}

class ChatInput extends React.Component {
  onSubmit(e) {
    new Chat(this.refs.chatInput.value).send()
    e.preventDefault()
  }
  render() {
    return (<form onSubmit={(e) => this.onSubmit(e)} ><input type="text" ref="chatInput"/></form>)
  }
}

class Main extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      text: "Hello World",
      messages: [],
    }

    this.chatEvent = eventbus.on('in:chat', (message) => {
      let { messages } = this.state
      messages.push(message)
      this.setState({
        messages: messages
      })
    })
  }
  componentWillUnmount() {
    this.chatEvent.off();
  }
  render() {
    let { messages } = this.state
    console.dir(messages)
    messages = messages.map((v, k) => <li key={"message-" + k}>{v}</li>)
    return (<div>
      {this.state.text}
      <ul className="messages">
        {messages}
      </ul>
      <ChatInput />
    </div>);
  }
}

ReactDOM.render(
  <Layout header={<h1>I AM GOD</h1>} left="Leftness" right="Rightness" main={<Main/>}/>,
  document.getElementById('application')
)



class Ping extends Message {
  constructor() {
    super('ping')
  }
}

class Chat extends Message {
  constructor(message) {
    super('chat')
    this.payload = message
  }
}

setInterval(function () {
  let msg = new Ping()
  msg.send()
}, 1000);


class DefaultMessageHandler extends MessageHandler {
  constructor() {
    super('default')
  }
  handle(data) {
    console.log("Default Handler caught message:", data)
  }
}
new DefaultMessageHandler();

class PingMessageHandler extends MessageHandler {
  constructor() {
    super('ping')
  }
  handle(data) {
    console.log('Ping:', data)
    eventbus.emit('in:ping', data)
  }
}
new PingMessageHandler();

class ChatMessageHandler extends MessageHandler {
  constructor() {
    super('chat')
  }
  handle(data) {
    eventbus.emit('in:chat', data.payload)
  }
}
new ChatMessageHandler();
