package main

import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
  // "github.com/huandu/facebook"
  "io/ioutil"
  "log"
  "os/exec"
)

type command interface {
  Execute()
}

// Ping Command
type pingCommand struct {
  Client *client
}

func newPingCommand(client *client, data interface{}) *pingCommand {
  return &pingCommand{
    Client: client,
  }
}

func (s *pingCommand) Execute() {
  s.Client.sendMessage(&clientOutMessage{
    Type:    "ping",
    Payload: "pong",
  })
}

type anacondaConfig struct {
  ConsumerKey    string
  ConsumerSecret string
  AccessToken    string
  AccessSecret   string
}

var twitterApi *anaconda.TwitterApi

func NewAnaconda() {
  twitter := anacondaConfig{
    config.Consumer.Key,
    config.Consumer.Secret,
    config.Access.Token,
    config.Access.Secret,
  }

  anaconda.SetConsumerKey(twitter.ConsumerKey)
  anaconda.SetConsumerSecret(twitter.ConsumerSecret)

  twitterApi = anaconda.NewTwitterApi(twitter.AccessToken, twitter.AccessSecret)
}


// Nick Command
type nickCommand struct {
  Client *client
  Nick   string
}

func newNickCommand(client *client, data interface{}) *nickCommand {
  return &nickCommand{
    Client: client,
    Nick:   data.(string),
  }

}

func (s *nickCommand) Execute() {
  old := s.Client.Name
  s.Client.Name = s.Nick
  go s.Client.room.broadcast(fmt.Sprintf("%s has changed their name to %s.", old, s.Client.Name))
}

// Nick Command
type scriptCommand struct {
  Client *client
  Script  string
  Name string
}

func newScriptCommand(client *client, data interface{}) *scriptCommand {
  payload := data.(map[string]interface {})
  return &scriptCommand{
    Client: client,
    Script:  payload["script"].(string),
    Name: payload["name"].(string),
  }
}

func (s *scriptCommand) Execute() {
  go s.Client.room.broadcast(fmt.Sprintf("%s has created script $%s", s.Client.Name, s.Name))
  go s.Client.room.broadcast(fmt.Sprintf("%s", s.Script))
  ioutil.WriteFile(fmt.Sprintf("jobs/%s.js", s.Name), []byte(s.Script), 0644)

}

// Nick Command
type runCommand struct {
  Client *client
  Name string
}

func newRunCommand(client *client, data interface{}) *runCommand {
  payload := data.(map[string]interface {})
  return &runCommand{
    Client: client,
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
    s.Client.room.broadcast(fmt.Sprintf("results: \n%s", out))
  }()
}

// Nick Command
type tweetSearchCommand struct {
  Client *client
  Query  string
}

func newSearchTwitterCommand(client *client, data interface{}) *tweetSearchCommand {
  return &tweetSearchCommand{
    Client: client,
    Query:  data.(string),
  }
}

func (s *tweetSearchCommand) Execute() {
  s.Client.room.sendAll(clientOutMessage{
    Type:    "chat",
    Payload: fmt.Sprintf("Searching Twitter: %s", s.Query),
  })
  searchResult, _ := twitterApi.GetSearch(s.Query, nil)
  for _, tweet := range searchResult.Statuses {
    s.Client.room.sendAll(clientOutMessage{
      Type:    "chat",
      Payload: fmt.Sprintf("Twitter Search Result: %s", tweet.Text),
    })
  }
}
// Nick Command
type tweetCommand struct {
  Client *client
  Tweet  string
}

func newTweetCommand(client *client, data interface{}) *tweetCommand {
  return &tweetCommand{
    Client: client,
    Tweet:  data.(string),
  }
}

func (s *tweetCommand) Execute() {
  twitterApi.PostTweet(s.Tweet, nil)
  fmt.Println("Tweet tweet!")
}

// Say Command
type sayCommand struct {
  Client *client
  Text   string
}

func newSayCommand(client *client, data interface{}) *sayCommand {
  return &sayCommand{
    Client: client,
    Text:   data.(string),
  }
}

func (s *sayCommand) Execute() {
  s.Client.room.sendAll(clientOutMessage{
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
