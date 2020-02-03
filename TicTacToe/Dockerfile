
#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git && apk update
WORKDIR $GOPATH/src/github.com/NehemiahG7/project-1/TicTacToe
COPY . .
RUN go build ./cmd/tictactoed

ENTRYPOINT ./tictactoed
LABEL Name=tictactoe Version=0.0.1
EXPOSE 8080
