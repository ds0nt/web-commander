import React from 'react'

import Events from '../lib/event-master'

import Room from './room'

export default class Rooms extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      rooms: ["home"],
      currentRoom: "home",
    }
    Events.on("in:join", this.addRoom.bind(this))
    Events.on("in:leave", this.removeRoom.bind(this))
    Events.on("action:tab:change", this.onTabChange.bind(this))
  }

  addRoom(roomName) {
    let rooms = this.state.rooms
    this.currentRoom = roomName
    rooms.push(roomName)
    this.setState({
      rooms: rooms
    })
  }

  removeRoom(roomName) {
    let rooms = this.state.rooms
    rooms = rooms.filter(r => r != roomName)
    this.setState({
      rooms: rooms
    })
  }

  onTabChange(roomName) {
    this.setState({currentRoom: roomName})
  }

  renderRoom(room, k) {
    return (
      <Room key={`room-${k}`} name={room} active={room == this.state.currentRoom}/>
    );
  }

  render() {
    let rooms = this.state.rooms.map(this.renderRoom.bind(this))

    return (<span>
      {rooms}
    </span>);
  }
}
