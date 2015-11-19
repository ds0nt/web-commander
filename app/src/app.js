import Layout from './layout'
import ReactDOM from 'react-dom'
import React from 'react'

class Main extends React.Component {
  constructor(props) {
    super(props)
    this.state = {
      text: "Hello World"
    }

    let i = 1;
    setInterval(() => {
      this.setState({text: "Hello World " + i})
      i++
    }, 1000);
  }
  render() {

    return (<div>{this.state.text}</div>);
  }
}

ReactDOM.render(
  <Layout
    header={<h1>I AM GOD</h1>}
    left="Leftness"
    right="Rightness"
    main={<Main/>}
    />, document.getElementById('application'))
