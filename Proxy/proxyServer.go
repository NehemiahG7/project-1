package main

import (
	"log"
	"fmt"
	"net"
)

func main(){
	ln, err := net.Listen("tcp", ":8080")
	fmt.Println("listening")

	if err != nil{
		log.Fatalf("listen err: %s", err)
	}
	for{
		conn, err := ln.Accept()
		defer conn.Close()
		fmt.Println("Accepted connection")
		if err != nil {
			log.Fatalf("conn err: %s", err)
		}
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		fmt.Printf("%s", buffer)
		
	}

}