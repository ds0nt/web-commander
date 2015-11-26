package main

import (

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

func (s *Switch) Run() {
  go func() {
    var msg *clientMessage
    for {
      msg = <-s.Messages
      var cmd command
      switch {
      case msg.Type == "ping":
        cmd = newPingCommand(msg.Client, msg.Payload)
      case msg.Type == "chat":
        cmd = newSayCommand(msg.Client, msg.Payload)
      case msg.Type == "nick":
        cmd = newNickCommand(msg.Client, msg.Payload)
      case msg.Type == "tweet":
        cmd = newTweetCommand(msg.Client, msg.Payload)
      default:
        cmd = newBadCommand(msg.Client)
      }
      s.Commands <-cmd
    }
  }()

  go func() {
    for {
      cmd := <- s.Commands
      cmd.Execute()
    }
  }()
}
