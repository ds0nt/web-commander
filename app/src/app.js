import Layout from './layout'



import ReactDOM from 'react-dom'
import React from 'react'


let addMessage = () => {}

class Main extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      text: "Hello World",
      messages: [],
    }

    let i = 1;
    setInterval(() => {
      this.setState({text: "Hello World " + i})
      i++
    }, 1000);
    addMessage = (message) => {
      let s = this.state
      s.messages.push(message)
      this.setState(s)
    }
  }
  render() {
    return (<div>
      {this.state.text}
      {this.state.messages.map((msg, k) => <span key={`message-${k}`}>{msg}</span>)}
    </div>);
  }
}

ReactDOM.render(
  <Layout header={<h1>I AM GOD</h1>} left="Leftness" right="Rightness" main={<Main/>}/>,
  document.getElementById('application')
)




import {Message, MessageHandler} from './socket-master'

class Ping extends Message {
  constructor() {
    super('ping')
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
    addMessage('Ping!')
  }
}
new DefaultMessageHandler();

class PingMessageHandler extends MessageHandler {
  constructor() {
    super('ping')
  }
  handle(data) {
    console.log('Ping:', data)

  }
}
