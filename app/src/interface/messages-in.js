import socket from '../lib/socket-master'
import Events from '../lib/event-master'

socket.register('default', {
  handle: data => {
    console.log('Default Handler caught message:', data)
    Events.emit('in:default', data)
  }
})
socket.register('ping', {
  handle: data => Events.roomEmit(data.room, 'in:ping', data)
})
socket.register('nick', {
  handle: data => Events.roomEmit(data.room, 'in:nick', data.payload)
})
socket.register('chat', {
  handle: data => Events.roomEmit(data.room, 'in:chat', data.payload)
})
socket.register('broadcast', {
  handle: data => Events.roomEmit(data.room, 'in:broadcast', data.payload)
})
