import Layout from './layout'
import eventbus from './event-master'
import { Message, MessageHandler } from './socket-master'

import ReactDOM from 'react-dom'
import React from 'react'


let addMessage = () => {}

class ChatInput extends React.Component {
  onSubmit(e) {
    e.preventDefault()
    let val = this.refs.chatInput.value
    this.refs.chatInput.value = ""
    if (val.startsWith('/nick')) {
      new Nick(val).send()
    } else if (val.startsWith('/tweet')) {
      new Tweet(val).send()
    } else {
      new Chat(val).send()
    }
  }
  render() {
    return (<form onSubmit={(e) => this.onSubmit(e)} ><input className="ui chat input" type="text" ref="chatInput"/></form>)
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

    this.chatEvent2 = eventbus.on('in:nick', (message) => {
      let { messages } = this.state
      messages.push(message)
      this.setState({
        messages: messages
      })
    })
  }
  componentWillUnmount() {
    this.chatEvent.off();
    this.chatEvent2.off();
  }
  render() {
    let { messages } = this.state
    messages = messages.map((v, k) => <div className="ui chat item" key={"message-" + k}>{v}</div>)
    return (
      <div className="ui list messages main-pane">
        <div className="main-pane-inner">
          <div className="ui header item main-pane-item">{this.state.text}</div>
          {messages}
          <ChatInput />
        </div>
      </div>
    );
  }
}

class HeaderPane extends React.Component {
  render() {
    let tabs = ["Room 1"]
    tabs = tabs.map(v => (<div className="ui item">{v}</div>))
    return (<div className="ui menu header-pane">
      <div className="ui header item header-pane-item">Forky</div>
      {tabs}
    </div>);
  }
}

class ChatPane extends React.Component {
  render() {
    let chatData = ["ds0nt", "john-bot", "jack-bot", "jill-bot"]
    let chatters = chatData.map(v => (<div className="item left-pane-item">{v}</div>))
    return (<div className="ui attached vertical menu left-pane">
      <div className="ui small header item">Room</div>
      {chatters}
    </div>);
  }
}

class AppPane extends React.Component {
  render() {
    let apps = ["abcd", "bcde", "cdef"]

    let cards = apps.map(v => (<div className="ui column right-pane-item">
    <div className="ui raised segment">
      <a className="ui black ribbon label">Overview</a>
      <span>{v}</span>
      <p></p>
    </div>
  </div>))
    return (<div className="ui two column padded grid right-pane">
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

class Nick extends Message {
  constructor(message) {
    super('nick')
    this.payload = message.slice(6)
  }
}

class Tweet extends Message {
  constructor(message) {
    super('tweet')
    this.payload = message.slice(7)
  }
}

setInterval(function () {
  let msg = new Ping()
  msg.send()
}, 10000);


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
    console.log("Received Message:", data)
    eventbus.emit('in:ping', data)
  }
}
new PingMessageHandler();

class NickMessageHandler extends MessageHandler {
  constructor() {
    super('nick')
  }
  handle(data) {
    console.log("Received Message:", data)
    eventbus.emit('in:nick', data.payloadcd)
  }
}
new NickMessageHandler();

class ChatMessageHandler extends MessageHandler {
  constructor() {
    super('chat')
  }
  handle(data) {
    console.log("Received Message:", data)
    eventbus.emit('in:chat', data.payload)
  }
}
new ChatMessageHandler();
