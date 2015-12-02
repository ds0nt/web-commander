import React from 'react'
import Events from '../lib/event-master'

export default class HeaderPane extends React.Component {
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

  changeTab(tab) {
    Events.emit("action:tab:change", tab)
  }

  onTabChange(roomName) {
    this.setState({currentRoom: roomName})
  }

  renderTab(tab, k) {
    let classes = "item"
    if (this.state.currentRoom == tab) {
      classes = `${classes} active`
    }
    return (
      <a key={`room-tab-${k}`} className={classes} onClick={e => this.changeTab(tab)}>
        {tab}
      </a>
    );
  }

  render() {
    let tabs = this.state.rooms.map(this.renderTab.bind(this))

    return (<div className="ui menu header-pane">
      <div className="ui header item header-pane-item">
        Forky
      </div>
      {tabs}
    </div>);
  }
}
