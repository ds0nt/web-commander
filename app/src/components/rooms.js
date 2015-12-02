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
    Events.on("in:join", this.addRoom)
    Events.on("in:leave", this.removeRoom)
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

  renderTab(tab, k) {
    return (
      <div key={`room-tab-${k}`} className="ui item" onClick={e => this.onTabChange(e)}>
        {tab}
      </div>
    );
  }

  render() {
    let tabs = this.state.rooms.map(this.renderTab.bind(this))
    let rooms = this.state.rooms.map(this.renderRoom.bind(this))

    return (<span>
      {tabs}
      {rooms}
    </span>);
  }
}
