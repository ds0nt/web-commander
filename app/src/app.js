import Layout from './layout'
import eventbus from './event-master'
import { Message, MessageHandler } from './socket-master'

import ReactDOM from 'react-dom'
import React from 'react'


let addMessage = () => {}

class ChatInput extends React.Component {
  onSubmit(e) {
    new Chat(this.refs.chatInput.value).send()
    this.refs.chatInput.value = ""
    e.preventDefault()
  }
  render() {
    return (<form onSubmit={(e) => this.onSubmit(e)} ><input className="ui input" type="text" ref="chatInput"/></form>)
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
    messages = messages.map((v, k) => <div className="ui chat item" key={"message-" + k}>ds0nt: {v}</div>)
    return (
      <div className="ui list messages">
        <div className="ui header item">{this.state.text}</div>
        {messages}
        <ChatInput />
      </div>
    );
  }
}

class HeaderPane extends React.Component {
  render() {
    return (<div className="ui menu"><div className="ui header item">Forky</div></div>);
  }
}

class ChatPane extends React.Component {
  render() {
    let chatData = ["ds0nt", "john-bot", "jack-bot", "jill-bot"]
    let chatters = chatData.map(v => (<div className="item">{v}</div>))
    return (<div className="ui inverted red attached vertical menu">
      <div className="ui small header item">Room</div>
      {chatters}
    </div>);
  }
}

class AppPane extends React.Component {
  render() {
    let apps = ["abcd", "bcde", "cdef"]

    let cards = apps.map(v => (<div className="ui inverted red column">
    <div className="ui raised segment">
      <a className="ui black ribbon label">Overview</a>
      <span>{v}</span>
      <p></p>
    </div>
  </div>))
    return (<div className="ui two column padded grid">
        {cards}
    </div>);
  }
}

ReactDOM.render(
  <Layout header={<HeaderPane />} left=<ChatPane /> right={<AppPane />} main={<Main/>}/>,
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
