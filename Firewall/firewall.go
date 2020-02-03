package main

import (
	"io"
	"net"
	"net/url"
	"net/http"
	"log"
	"fmt"
	"github.com/NehemiahG7/project-1/Firewall/Config"
)

//Channel to pass messages for the logger
var logCh chan string = make(chan string)

func main(){

	//Connect to logger
	go handleLog()

	proxy := http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request){
		//log
		logCh <- "request from " + req.RemoteAddr

		//create target
		tURL, err := url.Parse("http://127.0.0.1:" + config.LoadPort)
		if err != nil{
			log.Fatalf("tURL: %s\n", err)
		}

		//Copy to forward to target
		req.Host = tURL.Host
		req.URL.Host = tURL.Host
		req.URL.Scheme = tURL.Scheme
		req.RequestURI = ""

		// //get port from request
		// s, _, _ := net.SplitHostPort(req.RemoteAddr)
		// req.Header.Set("X-Forwarded-For", s)

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
	fmt.Printf("Proxy listening on port %s\n", config.FirewallPort)
	http.ListenAndServe(":" + config.FirewallPort, proxy)
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