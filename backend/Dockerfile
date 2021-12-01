FROM golang:1.17-alpine

RUN apk add git python3

WORKDIR /go/src/app

COPY . .

RUN go get -d -v ./...
RUN go get github.com/githubnemo/CompileDaemon
RUN go install -v ./...

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go install -v ./..." --command=run