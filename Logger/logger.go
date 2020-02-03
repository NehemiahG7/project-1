package main

import (
	"fmt"
	"net"
	"log"
	"os"
)

var fileName string = "txt.log"
var loggerPort string = ":9090"

func main(){
	file := loadFile(fileName)
	logger := log.New(file, loggerPort, log.Flags())
	defer file.Close()

	ln, err := net.Listen("tcp", loggerPort)
	defer ln.Close()

	if err != nil{
		logger.Fatalf("Listener failed: %s\n", err)
	}
	fmt.Printf("Logger listening on port :%s\n",loggerPort)

	connect := make(chan string)

	for{
		go handleLog(ln, logger, connect)
		logger.Printf("Connection made: %s\n", <-connect)
	}
}
func handleLog(ln net.Listener, logger *log.Logger, connect chan string){
	for{
		conn, err := ln.Accept()
		defer conn.Close()
		if err != nil{
			logger.Printf("Connection error: %s\n", err)
			return
		}

		logger.SetPrefix(conn.RemoteAddr().String() + " - ")

		connect <- string("connected")
		buff := make([]byte, 1024)

		num, err := conn.Read(buff)
		if err != nil{
			logger.Printf("Read failed: %s\n", err)
		}
		trm := buff[:num]

		logger.Printf("%s\n", trm)

	}
}
func loadFile(name string) *os.File{
	file, err := os.OpenFile(name, os.O_RDWR| os.O_APPEND| os.O_CREATE, 0666)
	if err != nil{
		log.Fatalf("Cannot open file: %s\n", err)
	}
	return file
}