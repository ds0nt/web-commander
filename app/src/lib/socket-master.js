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
      room: message.room,
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
    eventbus.emit('socket:disconnect')
  }

  register(type, messageHandler) {
    this.handlers[type] = messageHandler;
  }
}

export default new SocketMaster()
