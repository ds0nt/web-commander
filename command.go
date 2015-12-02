package main

import (
  "fmt"
  // "github.com/huandu/facebook"
  "io/ioutil"
  log "github.com/Sirupsen/logrus"
  "os/exec"
)

type command interface {
  Execute()
}

// Ping Command
type pingCommand struct {
  Client *client
  Room *room
}

func newPingCommand(client *client, room *room, data interface{}) *pingCommand {
  return &pingCommand{client, room}
}

func (s *pingCommand) Execute() {
  s.Client.sendMessage(&clientOutMessage{
    Type:    "ping",
    Payload: "pong",
  })
}

// Nick Command
type joinCommand struct {
  Client *client
  Room *room
  RoomName string
  Rooms *rooms
}

func newJoinCommand(client *client, room *room, data interface{}, rooms *rooms) *joinCommand {
  return &joinCommand{
    Client: client,
    Room: room,
    RoomName: data.(string),
    Rooms: rooms,
  }
}

func (s *joinCommand) Execute() {
  room := s.Rooms.getRoom(s.RoomName)
  s.Rooms.joinClient(s.Client, room)
  s.Client.sendMessage(&clientOutMessage{
    Type:    "join",
    Room: s.Room.Name,
    Payload: s.RoomName,
  })
}

// Nick Command
type nickCommand struct {
  Client *client
  Room *room
  Nick   string
}

func newNickCommand(client *client, room *room, data interface{}) *nickCommand {
  return &nickCommand{
    Client: client,
    Room: room,
    Nick:   data.(string),
  }

}

func (s *nickCommand) Execute() {
  old := s.Client.Name
  s.Client.Name = s.Nick
  go s.Room.broadcast(fmt.Sprintf("%s has changed their name to %s.", old, s.Client.Name))
}

// Nick Command
type scriptCommand struct {
  Client *client
  Room *room
  Script  string
  Name string
}

func newScriptCommand(client *client, room *room, data interface{}) *scriptCommand {
  payload := data.(map[string]interface {})
  return &scriptCommand{
    Client: client,
    Room: room,
    Script:  payload["script"].(string),
    Name: payload["name"].(string),
  }
}

func (s *scriptCommand) Execute() {
  go s.Room.broadcast(fmt.Sprintf("%s has created script $%s", s.Client.Name, s.Name))
  go s.Room.broadcast(fmt.Sprintf("%s", s.Script))
  ioutil.WriteFile(fmt.Sprintf("jobs/%s.js", s.Name), []byte(s.Script), 0644)

}

// Nick Command
type runCommand struct {
  Client *client
  Room *room
  Name string
}

func newRunCommand(client *client, room *room, data interface{}) *runCommand {
  payload := data.(map[string]interface {})
  return &runCommand{
    Client: client,
    Room: room,
    Name: payload["name"].(string),
  }
}

func (s *runCommand) Execute() {
  go func() {
    out, err := exec.Command("./run-job.sh", s.Name).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Running Script: %s\n", out)
    s.Room.broadcast(fmt.Sprintf("results: \n%s", out))
  }()
}

// Nick Command
type tweetSearchCommand struct {
  Client *client
  Room *room
  Query  string
}

func newSearchTwitterCommand(client *client, room *room, data interface{}) *tweetSearchCommand {
  return &tweetSearchCommand{
    Client: client,
    Room: room,
    Query:  data.(string),
  }
}

func (s *tweetSearchCommand) Execute() {
  s.Room.sendAll(clientOutMessage{
    Type:    "chat",
    Payload: fmt.Sprintf("Searching Twitter: %s", s.Query),
  })
  searchResult, _ := twitterApi.GetSearch(s.Query, nil)
  for _, tweet := range searchResult.Statuses {
    s.Room.sendAll(clientOutMessage{
      Type:    "chat",
      Payload: fmt.Sprintf("Twitter Search Result: %s", tweet.Text),
    })
  }
}
// Nick Command
type tweetCommand struct {
  Client *client
  Room *room
  Tweet  string
}

func newTweetCommand(client *client, room *room, data interface{}) *tweetCommand {
  return &tweetCommand{
    Client: client,
    Room: room,
    Tweet:  data.(string),
  }
}

func (s *tweetCommand) Execute() {
  twitterApi.PostTweet(s.Tweet, nil)
  s.Room.sendAll(clientOutMessage{
    Type:    "chat",
    Payload: fmt.Sprintf("Posted Tweet: %s", s.Client.Name, s.Tweet),
  })
}

// Say Command
type sayCommand struct {
  Client *client
  Room *room
  Text   string
}

func newSayCommand(client *client, room *room, data interface{}) *sayCommand {
  return &sayCommand{
    Client: client,
    Room: room,
    Text:   data.(string),
  }
}

func (s *sayCommand) Execute() {
  s.Room.sendAll(clientOutMessage{
    Type:    "chat",
    Payload: fmt.Sprintf("%s: %s", s.Client.Name, s.Text),
  })
}

// Say Command
type broadcastCommand struct {
  Room *room
  Text   string
}

func newBroadcastCommand(room *room, data interface{}) *broadcastCommand {
  return &broadcastCommand{
    Room: room,
    Text: data.(string),
  }
}

func (s *broadcastCommand) Execute() {
  s.Room.sendAll(clientOutMessage{
    Type:    "broadcast",
    Payload: fmt.Sprintf("%s", s.Text),
  })
}

// Bad Command
type badCommand struct {
  Client *client
}

func newBadCommand(client *client) *badCommand {
  return &badCommand{client}
}

func (s *badCommand) Execute() {
  log.Println("Bad Command")
}
