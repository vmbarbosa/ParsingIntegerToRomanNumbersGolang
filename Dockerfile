FROM golang:1.15

WORKDIR /go/src/integertoroman/backend

COPY [".", "/go/src/integertoroman"]

RUN go get -u github.com/gorilla/mux

RUN go get github.com/githubnemo/CompileDaemon

RUN ls

EXPOSE 3000

RUN CompileDaemon -command="./backend"
