package main

import (
  "net/http"
  "log"

  "github.com/rs/cors"
)

var cmds *commander

func main() {
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedHeaders: []string{"Sec-WebSocket-Extensions", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "Host", "X-Real-IP", "X-Forwarded-For", "X-Forwarded-Host"},
    AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
  })

  r := newRoom()
  cmds = newCommander()
  go cmds.Run()

  fs := http.FileServer(http.Dir("app/dist"))
  http.Handle("/", fs)
  http.Handle("/room", c.Handler(r))
  http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./app/dist/"))))
  // get the room going
  go r.run()

  // start the web server
  if err := http.ListenAndServe(":9080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
