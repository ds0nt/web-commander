package main

import (
  "net/http"
  "log"

  "github.com/rs/cors"
)

func main() {
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"*"},
    AllowedHeaders: []string{"Sec-WebSocket-Extensions", "Sec-WebSocket-Key", "Sec-WebSocket-Version", "Host", "X-Real-IP", "X-Forwarded-For", "X-Forwarded-Host"},
    AllowedMethods: []string{"GET", "POST", "OPTIONS", "PUT"},
  })

  r := newRoom()

  http.Handle("/", c.Handler(&templateHandler{filename: "chat.html"}))
  http.Handle("/room", c.Handler(r))

  // get the room going
  go r.run()

  // start the web server
  if err := http.ListenAndServe(":9080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
