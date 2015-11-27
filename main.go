package main

import (
  "flag"
  "log"
  "net/http"

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
  newRedis()
  NewAnaconda()

  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedHeaders: []string{"Sec-WebSocket-Extensions", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "Host", "X-Real-IP", "X-Forwarded-For", "X-Forwarded-Host"},
    AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
  })

  r = newRoom()
  commandSwitch = NewSwitch()

  go r.run()
  go commandSwitch.Run()

  fs := http.FileServer(http.Dir("app/dist"))
  http.Handle("/", fs)
  http.Handle("/room", c.Handler(r))
  http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./app/dist/"))))
  // get the room going

  // start the web server
  if err := http.ListenAndServe(":9080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
