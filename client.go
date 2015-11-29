package main

import (
  "github.com/gorilla/websocket"
  "log"
  "encoding/json"
)

type client struct {
  socket *websocket.Conn
  Name   string
  out   chan *clientOutMessage
  room   *room
}

type clientMessage struct {
  Type string `json:"type"`
  Payload interface{} `json:"payload"`
  Client *client
}
type clientOutMessage struct {
  Type string `json:"type"`
  Payload interface{} `json:"payload"`
}

func newClient(socket *websocket.Conn, room *room) *client {
  return &client{socket, "", make(chan *clientOutMessage, messageBufferSize), room }
}

func (c *client) sendMessage(cmsg *clientOutMessage) {
  c.out <- cmsg
}

func (c *client) doReadMessage() (*clientMessage, error) {
  var f clientMessage
  _, msg, err := c.socket.ReadMessage()
  if err != nil {
    return nil, err
  }
  err = json.Unmarshal(msg, &f)
  if err != nil {
    log.Printf("Bad JSON from client %s: %s", c.Name, string(msg))
    return nil, err
  }
  f.Client = c
  return &f, nil
}

func (c *client) doDrop() {
}

func (c *client) read() {
  for {
    msg, err := c.doReadMessage()
    if err != nil {
      log.Printf("Read Message Error %s", err)
      break
    }
    commandSwitch.Messages <-msg
  }

  log.Printf("Dropping Client %s: %s", c.Name)
  c.socket.Close()
  close(c.out)
  c.room.doLeave(c)
}

func (c *client) write() {
  for msg := range c.out {
    bytes, err := json.Marshal(&msg)
    if err != nil {
      log.Printf("Client Write Json Marshal Error: %v, %v", err, msg)
    }
    if err := c.socket.WriteMessage(websocket.TextMessage, bytes); err != nil {
      log.Printf("Client Write Error: %v", err)
      break
    }

    log.Printf("Client Send: %v", msg)
  }
}
