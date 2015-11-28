import Layout from './layout'
import eventbus from './event-master'

import notifications from './notifications'
notifications.start()

import { messageSwitch, Message, MessageHandler } from './socket-master'

import ReactDOM from 'react-dom'
import React from 'react'

require('codemirror/mode/javascript/javascript')
require('codemirror/mode/css/css')
require('codemirror/mode/yaml/yaml')

import codebox from 'codemirror'

export let scriptbox = null
class CodeBox extends React.Component {
  constructor(props) {
    super(props)
  }
  componentDidMount() {
    scriptbox = codebox.fromTextArea(this.refs.code, {
      mode: 'javascript',
      theme: 'monokai',
      inputStyle: "contenteditable",
      lineNumbers: true,
      tabsize: 2
    })
    this.setState({
      codeBox: scriptbox
    })
  }
  render() {
    return (<textarea ref="code"></textarea>)
  }

}

class ChatInput extends React.Component {
  onSubmit(e) {
    e.preventDefault()
    let val = this.refs.chatInput.value
    this.refs.chatInput.value = ""
    messageSwitch(val)
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

    this.chatEvent = eventbus.on('in:broadcast', (message) => this.message(message))
    this.chatEvent2 = eventbus.on('in:chat', (message) => this.message(message))
    this.chatEvent3 = eventbus.on('in:nick', (message) => this.message(message))
    this.chatEvent4 = eventbus.on('log', (message) => this.message(message))

  }

  message(message) {
    let { messages } = this.state
    messages.push(message)
    this.setState({
      messages: messages
    })
  }

  componentWillUnmount() {
    this.chatEvent.off();
    this.chatEvent2.off();
    this.chatEvent3.off();
    this.chatEvent4.off();
  }
  render() {
    let { messages } = this.state
    messages = messages.map((v, k) => <div className="ui chat item" key={"message-" + k}>{v}</div>)
    return (
      <div className="ui list messages main-pane">
        <div className="main-pane-inner">
          <div className="ui header item main-pane-item">{this.state.text}</div>
          {messages}
          <CodeBox />
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
