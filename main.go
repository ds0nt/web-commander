package main

import (
  "net/http"
  "log"
)
func main() {

  r := newRoom()

  http.Handle("/", &templateHandler{filename: "chat.html"})
  http.Handle("/room", r)

  // get the room going
  go r.run()

  // start the web server
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
