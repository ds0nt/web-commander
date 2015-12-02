import socket from '../lib/socket-master'

export default class MessageFactory {
  constructor(room) {
    this.room = room
  }
  send(txt) {
    let [_, lead, command, args] = txt.match(/(.)(\S*)(.*)/)
    args = args.substring(1)
    switch (lead) {

    // Slash Messages
    case '/':
      switch (command) {
      case 'join':
        return messages.join(this.room, args)
      case 'nick':
        return messages.nick(this.room, args)
      case 'cout':
        return messages.chat(this.room, "args\n" + scriptBoxValue())
      case 'tweet':
        return messages.tweet(this.room, args)
      case 'search-twitter':
        return messages.searchTwitter(this.room, args)
      default:
        return console.log('invalid command')
      }
      break;

    // Creation Messages
    case '$':
      return messages.script(this.room, command, scriptBoxValue())

    // Execution Messages
    case '!':
      return messages.run(this.room, command)
    }
    return messages.chat(this.room, txt)
  }
}

/// MESSAGES
let messages = {
  ping: (room) => socket.send({
    type: 'ping',
    room: room,
  }),
  chat: (room, txt) => socket.send({
    type: 'chat',
    room: room,
    payload: txt,
  }),
  nick: (room, txt) => socket.send({
    type: 'nick',
    room: room,
    payload: txt,
  }),
  tweet: (room, txt) => socket.send({
    type: 'tweet',
    room: room,
    payload: txt,
  }),
  searchTwitter: (room, txt) => socket.send({
    type: 'search-twitter',
    room: room,
    payload: txt,
  }),
  script: (room, name, script) => socket.send({
    type: 'script',
    payload: {
      name: name,
      room: room,
      script: script,
    }
  }),
  join: (room, roomName) => socket.send({
    type: 'join',
    room: room,
    payload: roomName,
  }),
  run: (room, name) => socket.send({
    type: 'run',
    room: room,
    payload: { name }
  }),
}
