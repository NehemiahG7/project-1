package main

import (
	"io"
	"net"
	"net/url"
	"net/http"
	"log"
	"fmt"
	"github.com/NehemiahG7/project-1/Proxy/Redirector"
	"github.com/NehemiahG7/project-1/Proxy/Config"
)

//Channel to pass messages for the logger
var logCh chan string = make(chan string)

func main(){

	//Connect to logger
	go handleLog()
	logCh <- "Proxy Online"

	//Create a proxy handler to forward requests
	proxy := http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request){
		logCh <- "Proxy: request from " + req.RemoteAddr
		str := balancer.GetPort()
		tURL, err := url.Parse("http://172.18.0.3:" + str)
		if err != nil{
			logCh <- "Proxy: Cannot Parse URL http://127.0.0.1:" + str + err.Error()
			log.Fatalf("tURL: %s\n", err)
		}

		req.Host = tURL.Host
		req.URL.Host = tURL.Host
		req.URL.Scheme = tURL.Scheme
		req.RequestURI = ""

		s, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil{
			logCh <- "Proxy: Cannot SplitHostPort " + req.RemoteAddr
		}
		req.Header.Set("X-Forwarded-For", s)
	

		resp, err := http.DefaultClient.Do(req)
		if err != nil{
			logCh <- "Proxy: Request could not be processed, " + err.Error()
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
		logCh <- "Proxy: Forwarded to: " + s
	})
	fmt.Printf("Proxy listening on port %s\n", config.ProxyPort)
	http.ListenAndServe(":" + config.ProxyPort, proxy)
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