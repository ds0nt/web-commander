import React from 'react'

export default class ChatPane extends React.Component {
  render() {
    let chatData = ["ds0nt", "john-bot", "jack-bot", "jill-bot"]
    let chatters = chatData.map(v => (<div className="item left-pane-item">{v}</div>))
    return (<div className="ui attached vertical menu left-pane">
      <div className="ui small header item">Room</div>
      {chatters}
    </div>);
  }
}
