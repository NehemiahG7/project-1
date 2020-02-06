
#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git && apk update
WORKDIR $GOPATH/src/github.com/NehemiahG7/project-1/LoadBalancer
COPY . .
RUN go build ./loadBalancer.go

ENTRYPOINT ["./loadBalancer"]
CMD [""]
LABEL Name=loadbalancer Version=0.0.1
EXPOSE 8082
