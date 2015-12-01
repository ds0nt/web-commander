require('codemirror/mode/javascript/javascript')
require('codemirror/mode/css/css')
require('codemirror/mode/yaml/yaml')
import codebox from 'codemirror'

import React from 'react'

let scriptbox = null

export class Scriptbox extends React.Component {
  constructor(props) {
    super(props)
  }
  componentDidMount() {
    scriptbox = codebox.fromTextArea(this.refs.code, {
      mode: 'javascript',
      theme: 'monokai',
      inputStyle: "contenteditable",
      lineNumbers: true,
      tabsize: 2
    })
  }
  render() {
    return (<textarea ref="code"></textarea>)
  }
}

export let scriptBoxValue = () => scriptbox.getValue()
