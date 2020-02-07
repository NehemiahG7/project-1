
#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git && apk update
WORKDIR $GOPATH/src/github.com/NehemiahG7/project-1/Proxy
COPY . .
RUN go build ./proxyServer.go


ENTRYPOINT ["./proxyServer"]
CMD [""]
LABEL Name=proxyServer Version=0.0.1
EXPOSE 8081
