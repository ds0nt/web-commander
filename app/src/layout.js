import React from 'react'

const HEADER_HEIGHT = '50px'
const LEFT_WIDTH = '200px'
const RIGHT_WIDTH = '300px'

class HeaderPane extends React.Component {
  render() {
    let style = {
      position: 'absolute',
      top: 0,
      left: 0,
      right: 0,
      height: HEADER_HEIGHT,
    }
    return (<div style={style} className="header-section">
      {this.props.children}
    </div>);
  }
}

class LeftPane extends React.Component {
  render() {
    let style = {
      position: 'absolute',
      left: 0,
      top: this.props.layoutState.header ? HEADER_HEIGHT : 0,
      bottom: 0,
      width: LEFT_WIDTH,
    }
    return (<div style={style} className="left-section">
      {this.props.children}
    </div>);
  }
}


class RightPane extends React.Component {
  render() {
    let style = {
      position: 'absolute',
      right: 0,
      top: this.props.layoutState.header ? HEADER_HEIGHT : 0,
      bottom: 0,
      width: RIGHT_WIDTH,
    }
    return (<div style={style} className="right-section">
      {this.props.children}
    </div>);
  }
}

class MainPane extends React.Component {
  render() {
    let style = {
      position: 'absolute',
      top: this.props.layoutState.header ? HEADER_HEIGHT : 0,
      left: this.props.layoutState.left ? LEFT_WIDTH : 0,
      right: this.props.layoutState.right ? RIGHT_WIDTH : 0,
      bottom: 0,
    }

    return (<div style={style} className="main-section">
      {this.props.children}
    </div>);
  }
}

class Layout extends React.Component {
  constructor(props) {
    super(props)
  }
  // mountHeader(content) {
  //   let $s = this.state
  //   $s.content.header = content
  //   this.setState($s)
  // }
  // mountLeft(content) {
  //   let $s = this.state
  //   $s.content.left = content
  //   this.setState($s)
  // }
  // mountRight(content) {
  //   let $s = this.state
  //   $s.content.right = content
  //   this.setState($s)
  // }
  // mountMain(content) {
  //   let $s = this.state
  //   $s.content.main = content
  //   this.setState($s)
  // }
  render() {
    let children = [];
    let $p = this.props

    if ($p.layout.header)
      children.push(<HeaderPane layoutState={this.props.layout} key="headerPane">{$p.header}</HeaderPane>)
    if ($p.layout.main)
      children.push(<MainPane layoutState={this.props.layout} key="mainPane">{$p.main}</MainPane>)
    if ($p.layout.left)
      children.push(<LeftPane layoutState={this.props.layout} key="leftPane">{$p.left}</LeftPane>)
    if ($p.layout.right)
      children.push(<RightPane layoutState={this.props.layout} key="rightPane">{$p.right}</RightPane>)

    return (<div>{children}</div>)
  }
}
Layout.defaultProps = {
  layout: {
    header: true,
    main: true,
    left: true,
    right: true,
  },
  header: 'Top Header',
  main: 'Main Pane',
  left: 'Left Pane',
  right: 'Right Pane',
}
export default Layout
