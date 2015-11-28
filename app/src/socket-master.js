import scriptbox from './app'
import eventbus from './event-master'

const WEBSOCKET_ENDPOINT = `ws://${window.location.host}/room`

class SocketMaster {
  constructor() {
    this.socket = new WebSocket(WEBSOCKET_ENDPOINT)
    this.socket.onopen = this.onOpen.bind(this)
    this.socket.onmessage = this.onMessage.bind(this)
    this.socket.onclose = this.onClose.bind(this)
    this.socket.onerror = this.onError.bind(this)
    this.handlers = {}
    console.dir(this.socket)
  }

  send(message) {
    this.socket.send(JSON.stringify({
      type: message.type,
      payload: message.payload
    }))
  }

  onOpen(e) {
    console.log('Socket Master onOpen', e)
  }

  onMessage(message) {
    let data = JSON.parse(message.data)
    if (typeof this.handlers[data.type] === 'undefined') {
      this.handlers['default'].handle(data)
    } else {
      this.handlers[data.type].handle(data)
    }
  }
  onError(e) {
    console.log('Socket Master onError', e)
  }
  onClose(e) {
    console.log('Socket Master onClose', e)
  }

  register(type, messageHandler) {
    this.handlers[type] = messageHandler;
  }
}
var SocketMain = new SocketMaster()

function messageSwitch(txt) {
  let [_, lead, command, args] = txt.match(/(.)(\S*)(.*)/)

  switch (lead) {
  case '/':
    switch (command) {
    case 'nick':
      return Send.nick(args)
    case 'tweet':
      return Send.tweet(args)
    case 'search-twitter':
      return Send.searchTwitter(args)
    default:
      return console.log('invalid command')
    }
    break;
  case '$':
    return Send.script(command, scriptbox.getValue())
  case '!':
    return Send.run(command)
  }
  Send.chat(txt)
}


/// MESSAGES
let Send = {
  ping: () => SocketMain.send({
    type: 'ping'
  }),
  chat: txt => SocketMain.send({
    type: 'chat',
    payload: txt,
  }),
  nick: txt => SocketMain.send({
    type: 'nick',
    payload: txt,
  }),
  tweet: txt => SocketMain.send({
    type: 'tweet',
    payload: txt,
  }),
  searchTwitter: txt => SocketMain.send({
    type: 'search-twitter',
    payload: txt,
  }),
  script: (name, script) => SocketMain.send({
    type: 'script',
    payload: {
      name: name,
      script: script,
    }
  }),
  run: name => SocketMain.send({
    type: 'run',
    payload: { name }
  }),
}

SocketMain.register('default', {
  handle: data => {
    console.log('Default Handler caught message:', data)
    eventbus.emit('in:default', data)
  }
})
SocketMain.register('ping', {
  handle: data => eventbus.emit('in:ping', data)
})
SocketMain.register('nick', {
  handle: data => eventbus.emit('in:nick', data.payload)
})
SocketMain.register('chat', {
  handle: data => eventbus.emit('in:chat', data.payload)
})
SocketMain.register('broadcast', {
  handle: data => eventbus.emit('in:broadcast', data.payload)
})

export default { messageSwitch, Send, SocketMain }
