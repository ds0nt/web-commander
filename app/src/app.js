import React from 'react'
import ReactDOM from 'react-dom'
import notifications from './lib/notifications'
import messages from './interface/messages-in'

import Layout from './components/layout'
import HeaderPane from './components/header'
import UserList from './components/user-list'
import AssetsList from './components/assets-list'
import Rooms from './components/rooms'

notifications.start()

ReactDOM.render(
  <Layout
    header={<HeaderPane />}
    left={<UserList />}
    right={<AssetsList />}
    main={<Rooms/>}
    />,
  document.getElementById('application')
)
