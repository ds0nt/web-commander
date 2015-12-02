package main

import (
  // "github.com/gorilla/mux"
  log "github.com/Sirupsen/logrus"
  "net/http"
)

type rooms struct {
  rooms map[string]*room
  clientRooms map[*client][]*room
}

func newRooms() *rooms {
  return &rooms{
    rooms: make(map[string]*room),
    clientRooms: make(map[*client][]*room),
  }
}

func (r *rooms) getRoom(roomName string) *room {
  room, ok := r.rooms[roomName]
  if !ok {
    room = newRoom(roomName)
    r.rooms[roomName] = room
    go room.run()
    log.Printf("New Room: %s\n", roomName)
  } else {
    log.Printf("Lookup Room: %s\n", roomName)
  }
  return room
}

func (r *rooms) joinClient(client *client, room *room) {
  r.clientRooms[client] = append(r.clientRooms[client], room)
  room.doJoin(client)
}

func (r *rooms) dropClient(client *client) {
  for _, room := range r.clientRooms[client] {
    log.Printf("Client %s left room %s\n", client.Name, room.Name)
    room.doLeave(client)
  }
}


func (r *rooms) Handle(w http.ResponseWriter, req *http.Request) {
  log.Println("Connecting to room")
  // vars := mux.Vars(req)
  // roomId := vars["id"]

  socket, err := upgrader.Upgrade(w, req, nil)
  if err != nil {
    log.Printf("Web Socket Upgrade: : %s\n", err)
    return
  }
  log.Printf("Web Socket Upgraded")

  client := newClient(socket)
  client.Name = "anonymous"

  home := r.getRoom("home");
  r.joinClient(client, home)
  defer r.dropClient(client)

  go client.write()
  client.read()
}
