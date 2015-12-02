package main

import (
  log "github.com/Sirupsen/logrus"
)

type Switch struct {
  Messages chan *clientMessage
  Commands chan command
}

func NewSwitch() *Switch {
  return &Switch{
    make(chan *clientMessage),
    make(chan command),
  }
}

func (s *Switch) Run(rooms *rooms) {
  go func() {
    var msg *clientMessage
    for {
      msg = <-s.Messages
      log.Printf("Routing Message: %s", msg)
      room := rooms.getRoom(msg.Room)
      log.Printf("Client %s sent of type %s to room %s\n", msg.Client.Name, msg.Type, room)
      var cmd command
      switch {
      case msg.Type == "ping":
        cmd = newPingCommand(msg.Client, room, msg.Payload)
      case msg.Type == "join":
        cmd = newJoinCommand(msg.Client, room, msg.Payload, rooms)
      case msg.Type == "chat":
        cmd = newSayCommand(msg.Client, room, msg.Payload)
      case msg.Type == "nick":
        cmd = newNickCommand(msg.Client, room, msg.Payload)
      case msg.Type == "tweet":
        cmd = newTweetCommand(msg.Client, room, msg.Payload)
      case msg.Type == "script":
        cmd = newScriptCommand(msg.Client, room, msg.Payload)
      case msg.Type == "run":
        cmd = newRunCommand(msg.Client, room, msg.Payload)
      case msg.Type == "search-twitter":
        cmd = newSearchTwitterCommand(msg.Client, room, msg.Payload)
      default:
        cmd = newBadCommand(msg.Client)
      }
      if cmd != nil {
        s.Commands <-cmd
      }
    }
  }()

  go func() {
    for {
      cmd := <- s.Commands
      cmd.Execute()
    }
  }()
}
