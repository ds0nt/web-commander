package main

import (
  "github.com/gorilla/websocket"
  "log"
  "encoding/json"
)

type client struct {
  socket *websocket.Conn
  Name   string
  send   chan clientOutMessage
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
  return &client{socket, "", make(chan clientOutMessage, messageBufferSize), room }
}

func (c *client) read() {
  for {
    _, msg, err := c.socket.ReadMessage()
    if err == nil {
      var f clientMessage
      err := json.Unmarshal(msg, &f)
      f.Client = c
      if err != nil {
        log.Printf("Evil JSON Detected: %v, %v", err, string(msg))
        continue
      }
      commandSwitch.Messages <- &f
    } else {
      break
    }
  }
  c.socket.Close()
}

func (c *client) write() {
  for msg := range c.send {
    bytes, err := json.Marshal(&msg)
    if err != nil {
      log.Printf("Client Write Json Marshal Error: %v, %v", err, msg)
    }
    if err := c.socket.WriteMessage(websocket.TextMessage, bytes); err != nil {
      log.Printf("Client Write Error: %v", err)
      break
    } else {
      log.Printf("Client Send: %v", msg)
    }
  }
}
