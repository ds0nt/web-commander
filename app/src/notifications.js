import Notify from 'notifyjs'
import events from './event-master'

function doNotification(title, text) {
  var myNotification = new Notify(title, {
      body: text,
      notifyShow: () => console.log('notification was shown!'),
      timeout: 3,
  })

  if (!Notify.needsPermission) {
    myNotification.show()
  } else if (Notify.isSupported()) {
      Notify.requestPermission(onPermissionGranted, onPermissionDenied)
  }

  function onPermissionGranted() {
      console.log('Permission has been granted by the user')
      myNotification.show()
  }

  function onPermissionDenied() {
      console.warn('Permission has been denied by the user')
  }
}

export default {
  start() {
      events.on('in:chat', (txt) => {
        doNotification(window.location.hostname, txt)
      })
  }
}
