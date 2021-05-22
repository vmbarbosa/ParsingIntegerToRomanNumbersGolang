FROM golang:1.15

WORKDIR /go/src/

COPY [".", "/go/src/"]

RUN go get -u github.com/gorilla/mux

RUN go get github.com/githubnemo/CompileDaemon

RUN pwd

RUN ls

RUN cd backend/

RUN ls

RUN CompileDaemon 

EXPOSE 3000
