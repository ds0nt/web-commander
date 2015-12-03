import React from 'react'
import ReactDOM from 'react-dom'

export default class ThreeTest extends React.Component {
  constructor(props) {
    super(props)

    this.camera = new THREE.PerspectiveCamera( 75, window.innerWidth / window.innerHeight, 1, 10000 );
    this.camera.position.z = 1000;

    this.scene = new THREE.Scene();

    this.geometry = new THREE.BoxGeometry( 200, 200, 200 );
    this.material = new THREE.MeshBasicMaterial( { color: 0xff0000, wireframe: true } );

    this.mesh = new THREE.Mesh( this.geometry, this.material );
    this.scene.add( this.mesh );

    this.renderer = new THREE.WebGLRenderer({ antialias: true });
    this.renderer.setSize( window.innerWidth, window.innerHeight );
  }
  componentDidMount() {
    ReactDOM.findDOMNode(this.refs.t).appendChild(this.renderer.domElement)
    this.animate()
  }
  animate() {
     // note: three.js includes requestAnimationFrame shim
     requestAnimationFrame( this.animate );

     this.mesh.rotation.x += 0.01;
     this.mesh.rotation.y += 0.02;

     this.renderer.render( this.scene, this.camera );

  }

  render() {
    return (
      <div ref="t" />
    );
  }
}
