import { EventEmitter } from 'events'



class Events extends EventEmitter {
  constructor() {
    super()
  }
  roomEmit(room, ev, ...x) {
    return this.emit(`${room}:${ev}`, ...x)
  }
  roomOn(room, ev, ...x) {
    return this.on(`${room}:${ev}`, ...x)
  }
}

export default new Events();
