FROM golang:1.15

WORKDIR /usr/src

COPY [".", "/usr/src/"]

RUN go get -u github.com/gorilla/mux

RUN go get github.com/githubnemo/CompileDaemon

RUN CompileDaemon -command="./backend"

EXPOSE 3000
