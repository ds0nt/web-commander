FROM ubuntu

EXPOSE 9080

WORKDIR /src
COPY . /src/

CMD ./web-commander

