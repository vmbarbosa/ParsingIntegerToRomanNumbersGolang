FROM golang:1.15

WORKDIR /go/src/integertoroman/backend

COPY [".", "/go/src/integertoroman"]

RUN go get -u github.com/gorilla/mux

RUN go get github.com/githubnemo/CompileDaemon

RUN pwd

RUN ls

RUN cd backend/

RUN ls

RUN CompileDaemon 

EXPOSE 3000
