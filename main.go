package main

import (
  "flag"
  log "github.com/Sirupsen/logrus"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
  "github.com/ChimeraCoder/anaconda"
)

var Rooms *rooms
var commandSwitch *Switch
var (
  conf = flag.String("conf", "config.toml", "path to toml config")
)

func main() {
  flag.Parse()

  // la config
  loadConfig()

  // la commented commented redis
  // newRedis()

  // la twitter
  NewAnaconda()

  // la condoms
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedHeaders: []string{"Sec-WebSocket-Extensions", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "Host", "X-Real-IP", "X-Forwarded-For", "X-Forwarded-Host"},
    AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
  })

  // la room master
  rooms := newRooms()

  // la websocket message master
  commandSwitch = NewSwitch()
  go commandSwitch.Run(rooms)


  // la http router
  router := mux.NewRouter()
  fs := http.FileServer(http.Dir("app/dist"))

  // les routes
  router.HandleFunc("/room", rooms.Handle)
  router.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./app/dist/"))))
  router.PathPrefix("/").Handler(fs)
  http.Handle("/", c.Handler(router))
  // get the room going

  // la main screen
  if err := http.ListenAndServe(":9080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
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
