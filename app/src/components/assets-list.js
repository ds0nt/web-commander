import React from 'react'

export default class AppPane extends React.Component {
  render() {
    let apps = ["abcd", "bcde", "cdef"]

    let cards = apps.map(v => (<div className="ui column right-pane-item">
    <div className="ui raised segment">
      <a className="ui black ribbon label">Overview</a>
      <span>{v}</span>
      <p></p>
    </div>
  </div>))
    return (<div className="ui two column padded grid right-pane">
        {cards}
    </div>);
  }
}
