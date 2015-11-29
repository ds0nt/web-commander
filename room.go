package main

import (
  "github.com/gorilla/websocket"
  "fmt"
  "log"
  "net/http"
)

type room struct {
  counter int
  forward chan clientOutMessage
  join chan *client
  leave chan *client
  commands chan clientOutMessage
  clients map[*client]bool
}

type roomMessage struct {
  Type string `json:"type"`
  Payload interface{} `json:"payload"`
}


func newRoom() *room {
  room := &room{
    counter: 0,
    forward: make(chan clientOutMessage),
    join:    make(chan *client),
    leave:   make(chan *client),
    clients: make(map[*client]bool),
  }
  log.Printf("Creating Room: %v", room)
  return room
}

func (r *room) broadcast(text string) {
  commandSwitch.Commands <-newBroadcastCommand(r, text)
}

func (r *room) sendAll(cmd clientOutMessage) {
  r.forward <- cmd
}

func (r *room) sendOne(client *client, cmd clientOutMessage) {
  go client.sendMessage(&cmd)
}

func (r *room) joinClient(client *client) {
  r.join <- client
}

func (r *room) leaveClient(client *client) {
  r.leave <- client
}

func (r *room) doJoin(client *client) {
  r.clients[client] = true

  go r.broadcast(fmt.Sprintf("%s has joined the channel.", client.Name))
  client.sendMessage(&clientOutMessage{"chat", "Welcome to web commander."})
  client.sendMessage(&clientOutMessage{"chat", "The current list of users are:"})

  for c := range r.clients {
    client.sendMessage(&clientOutMessage{"chat", c.Name})
  }
}

func (r *room) doLeave(client *client) {
  go r.broadcast(fmt.Sprintf("%s has left the channel.", client.Name))
  delete(r.clients, client)
}

func (r *room) doSendAll(msg clientOutMessage) {
    for client := range r.clients {
      client.sendMessage(&msg)
    }
}

func (r *room) run() {
  for {
    select {
    case client := <-r.join:
      r.doJoin(client)
    case client := <-r.leave:
      r.doLeave(client)
    case msg := <-r.forward:
      r.doSendAll(msg)
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
  r.joinClient(client)
  defer r.leaveClient(client)

  go client.write()
  client.read()
}
