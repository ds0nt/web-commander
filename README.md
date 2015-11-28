# Forky Commander

Forky Commander is like Slack with Customizable Super Powers.


## Built-in Commands

```text

Change your nickname:
/nick ds0nt

Post on twitter:
/tweet forky commander is better then slack :) http://datahexagon:9080


Search twitter:
/search-twitter cute kittens


Print Code to Chat:
/cout

Save Command (see below):
$hello-world

Run Command:
!hello-world

```


## Creating your own commands.

In the code-box, enter in some simple nodejs code..
  
```javascript
var fs = require('fs')
console.log('hello world')
console.log((fs.readdirSync('/'))
```

In the chat box, using the ``$mycommand`` will create a command runnable as ``!mycommand``

```text
$hello-world
```

Now run the command:

```text
!hello-world
```

You should see the output:

```text
hello world
['bin','etc' .... ]
```

There is currently not support for npm requires... lol.. :) feel free to fork forky, and submit a pull request.


## Development

The chat server is written in Golang

  - Manages WebSocket Clients
  - Serves Browser Assets
  - Channels messages into rooms, and to clients
  - Has a common interface for commands, and a commandSwitch implementation that proxies messages into commands
  - Launches jobs in the ./jobs folder, using a script that runs a node docker image


The chat client is written in Reactjs / Browserify / Myth... it's pretty straightforward


## Roadmap (Golang)

- Add all sorts of built-in commands to support people's various professions

- Create a simple user system, to grant access to peoples authorized tokens for various third-party integrations.

- Documents, Scripts, and mindmaps that are within the application, for real-time collaboration hotness.

- Create a CLI interface, so that developers can pipe various logs into chat-rooms

- Copy the implementation from drone.io for managing docker jobs.

- Hook all rooms up to kafka, and create a RoomClient structure.

- Kafka will be used for a central messaging queue which will allow scaling of the socket server itself.

- Add simple data-stashes for user scripts

- Web Hook Powers (inbound messages from third-party)


## Front-end (ES6 Babelage)


- localstorage <- nicknames
- user list
- add a bunch of simple commands
  - /me
  - /slap
  - /join
  - /leave
  - /history
  - /grep

The architecture is built to be extensive. Built-in should be easy to add, on the server and client, it should be easy to extend chat messages and behaviours, and it should be easy to extend and customize docker jobs
