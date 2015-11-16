package main

import (
  "github.com/gorilla/websocket"
  "log"
)

type client struct {
  socket *websocket.Conn
  send   chan []byte
  room   *room
}


func (c *client) read() {
  for {
    _, msg, err := c.socket.ReadMessage()
    if err == nil {
      str := string(msg)
      if str[0:1] == "/" {
        cmd := NewCommand(c, c.room, str)
        cmds.commands <- cmd
      }
      c.room.forward <- msg
    } else {
      break
    }
  }
  c.socket.Close()
}

func (c *client) write() {
  for msg := range c.send {
    if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
      log.Printf("Client Write Error: %v", err)
      break
    } else {
      log.Println("Client Write")
    }
  }
}
