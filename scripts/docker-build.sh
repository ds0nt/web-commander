#1/bin/bash

cd "$(dirname $0)/.."

docker build -t "web-commander" .


  docker stop web-commander;
  docker rm web-commander;
  docker run -d -p 9080:9080 --name="web-commander" web-commander
  docker logs -f web-commander
