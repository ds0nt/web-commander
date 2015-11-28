package main

import (
  "github.com/gorilla/websocket"
  "fmt"
  "log"
  "net/http"
)

type room struct {
  counter int
  forward chan roomMessage
  join chan *client
  leave chan *client
  commands chan roomMessage
  clients map[*client]bool
}

type roomMessage struct {
  Type string `json:"type"`
  Payload interface{} `json:"payload"`
}


func newRoom() *room {
  room := &room{
    counter: 0,
    forward: make(chan roomMessage),
    join:    make(chan *client),
    leave:   make(chan *client),
    commands:   make(chan roomMessage),
    clients: make(map[*client]bool),
  }
  log.Printf("Creating Room: %v", room)
  return room
}

func (r *room) broadcast(text string) {
  commandSwitch.Commands <-newBroadcastCommand(r, text)
}

func (r *room) run() {
  for {
    select {
    case cmd := <-r.commands:
      r.forward <- cmd
    case client := <-r.join:
      r.clients[client] = true
      client.Name = fmt.Sprintf("anonymous%d", r.counter)
      r.counter++
      go r.broadcast(fmt.Sprintf("%s has joined the channel.", client.Name))
    case client := <-r.leave:
      go r.broadcast(fmt.Sprintf("%s has left the channel.", client.Name))
      delete(r.clients, client)
      close(client.send)
    case msg := <-r.forward:
      for client := range r.clients {
        out := clientOutMessage{
          Type: msg.Type,
          Payload: msg.Payload,
        }
        select {
        case client.send <-out:
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

  client := newClient(socket, r)
  r.join <- client
  defer func() {
    r.leave <- client
  }()
  go client.write()
  client.read()
}
