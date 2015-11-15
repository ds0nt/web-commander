package main

import (
  "github.com/gorilla/websocket"
  "log"
  "net/http"
)

type room struct {

  // forward is a channel that holds incoming messages
  // that should be forwarded to the other clients.
  forward chan []byte
  // join is a channel for clients wishing to join the room.
  join chan *client
  // leave is a channel for clients wishing to leave the room.
  leave chan *client
  // clients holds all current clients in this room.
  clients map[*client]bool
}

func newRoom() *room {
  room := &room{
    forward: make(chan []byte),
    join:    make(chan *client),
    leave:   make(chan *client),
    clients: make(map[*client]bool),
  }
  log.Printf("Creating Room: %v", room)
  return room
}

func (r *room) run() {
  for {
    select {
    case client := <-r.join:
      r.clients[client] = true
    case client := <-r.leave:
      delete(r.clients, client)
      close(client.send)
    case msg := <-r.forward:
      for client := range r.clients {
        select {
        case client.send <- msg:
        default:
          delete(r.clients, client)
          close(client.send)
        }
      }
    }
  }
}

const (
  socketBufferSize  = 1024
  messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{
  ReadBufferSize: socketBufferSize,
  WriteBufferSize: socketBufferSize,
  CheckOrigin: func(r *http.Request) bool {
    return true
  },
}

func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
  log.Printf("Web Socket Upgrade")
  socket, err := upgrader.Upgrade(w, req, nil)
  if err != nil {
    log.Fatal("ServeHTTP:", err)
    return
  }

  client := &client{
    socket: socket,
    send:   make(chan []byte, messageBufferSize),
    room:   r,
  }
  r.join <- client
  defer func() {
    r.leave <- client
  }()
  go client.write()
  client.read()
}
