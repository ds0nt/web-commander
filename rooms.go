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
    log.Printf("New Room %s\n", roomName)
    newRoom := newRoom(roomName)
    r.rooms[roomName] = newRoom
    go newRoom.run()
    log.Printf("newRoom: %s\n", newRoom)
    return newRoom
  }
  log.Printf("room: %s\n", room)
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
