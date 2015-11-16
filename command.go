package main

import (
  "log"
  "fmt"
)

type commander struct {
  commands chan Command
}

type Command interface {
  Execute()
}

type SayCommand struct {
  client *client
  room *room
  msg string
}
type BadCommand struct {}

func (s *SayCommand) Execute() {
  log.Println("SAAAY")
  s.room.forward <-[]byte(fmt.Sprintf("[s] %s\n", s.msg))
}

func (s *BadCommand) Execute() {
  log.Println("Bad Command")
}

func NewCommand(c *client, r *room, msg string) Command {
  log.Printf("Parsing Command: %v\n", msg)
  switch {
  case msg[0:4] == "/say":
    return &SayCommand{c, r, msg[4:]}
  default:
    return &BadCommand{}
  }
}

func (c *commander) Run() {
  log.Println("Starting Command Worker")
  for {
    cmd := <-c.commands
    log.Println("Executing Command")
    cmd.Execute()
  }
}

func newCommander() *commander {
  commander := &commander{
    commands: make(chan Command),
  }
  log.Printf("Creating commander: %v", commander)
  return commander
}
