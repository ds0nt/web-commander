package main

import (
  "flag"
  "log"
  "net/http"
  "github.com/gorilla/mux"
  "github.com/rs/cors"
)

var r *room
var commandSwitch *Switch
var (
  conf = flag.String("conf", "config.toml", "path to toml config")
)

func main() {
  flag.Parse()
  loadConfig()
  // newRedis()
  NewAnaconda()

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedHeaders: []string{"Sec-WebSocket-Extensions", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "Host", "X-Real-IP", "X-Forwarded-For", "X-Forwarded-Host"},
    AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
  })

  commandSwitch = NewSwitch()
  go commandSwitch.Run()

  rooms := newRooms()

  router := mux.NewRouter()
  fs := http.FileServer(http.Dir("app/dist"))


  router.HandleFunc("/room/{roomId}", rooms.Handle)
  router.PathPrefix("/assets").Handler(http.StripPrefix("/assets", http.FileServer(http.Dir("./app/dist/"))))
  router.PathPrefix("/").Handler(fs)
  http.Handle("/", c.Handler(router))
  // get the room going

  // start the web server
  if err := http.ListenAndServe(":9080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
