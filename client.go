package main

import (
  "github.com/gorilla/websocket"
  "log"
  "encoding/json"
)

type client struct {
  socket *websocket.Conn
  send   chan []byte
  room   *room
}

type message struct {
  Type string `json:"type"`
  Payload interface{} `json:"payload"`
}

func (c *client) read() {
  for {
    _, msg, err := c.socket.ReadMessage()
    if err == nil {
      var f message
      err := json.Unmarshal(msg, &f)
      if err != nil {
        log.Printf("Evil JSON Detected: %v, %v", err, string(msg))
        continue
      }
      switch f.Type {
      case "chat":
        c.room.forward <- []byte(f.Payload.(string))
      case "command":
        cmd := NewCommand(c, c.room, f.Payload.(string))
        cmds.commands <- cmd
      }

    } else {
      break
    }
  }
  c.socket.Close()
}

func (c *client) write() {
  for msg := range c.send {
    var f message
    f.Type = "chat"
    f.Payload = string(msg)
    bytes, err := json.Marshal(&f)
    if err != nil {
      log.Printf("Client Write Json Marshal Error: %v, %v", err, msg)
    }
    if err := c.socket.WriteMessage(websocket.TextMessage, bytes); err != nil {
      log.Printf("Client Write Error: %v", err)
      break
    } else {
      log.Println("Client Write")
    }
  }
}
