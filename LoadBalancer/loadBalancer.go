package main

import (
	"io"
	"fmt"
	"net"
	"log"
	"net/url"
	"net/http"
	"github.com/NehemiahG7/project-1/LoadBalancer/Config"
	"github.com/NehemiahG7/project-1/LoadBalancer/ServerType"

)

//Channel to pass messages for the logger
var logCh chan string = make(chan string)

//Servers is a copy of config.Servers
var Servers *[]serverhand.Server = config.Servers

func main(){

	//Connect to logger
	go handleLog()

	proxy := http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request){
		logCh <- "request from " + req.RemoteAddr
		tURL, err := url.Parse(serverhand.GetServer(*Servers))
		if err != nil{
			log.Fatalf("tURL: %s\n", err)
		}

		req.Host = tURL.Host
		req.URL.Host = tURL.Host
		req.URL.Scheme = tURL.Scheme
		req.RequestURI = ""

		s, _, _ := net.SplitHostPort(req.RemoteAddr)

		resp, err := http.DefaultClient.Do(req)
		if err != nil{
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(rw, err)
			return
		}

		for key, values := range resp.Header{
			for _,  value := range values{
				rw.Header().Set(key, value)
			}
		}
		rw.WriteHeader(resp.StatusCode)
		io.Copy(rw, resp.Body)
		logCh <- "Forwarded to: " + s
	})
	fmt.Printf("Balancer listening from docker net\n")
	http.ListenAndServe(":" + config.LoggerPort, proxy)
}
func handleLog(){
	for{
		conn, err := net.Dial("tcp", "127.0.0.1:" + config.LoggerPort)
		if err != nil{
			log.Fatalf("Dial errL %s\n", err)
		}
		conn.Write([]byte(<-logCh))
	}
	
}