FROM golang:latest
WORKDIR /go/src
RUN git clone https://github.com/jonathanschmittblog/jsapi-websocket.git
WORKDIR /go/src/jsapi-websocket
RUN go mod init
RUN go get -u
RUN go install
EXPOSE 3001