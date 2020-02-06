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
var Servers map[string]int = *config.Servers

func main(){

	//Connect to logger
	go handleLog()
	logCh <- "Balancer Online"

	logCh <- "Balancer: Serving: " + string(serverhand.GetKeys(Servers))
	proxy := http.HandlerFunc(func (rw http.ResponseWriter, req *http.Request){

		logCh <- "Balancer: request recieved from proxy for: " + req.Header["X-Forwarded-For"][0]


		se := serverhand.GetServer(Servers)
		Servers[se]++
		//Use this command when running loadBalancer outside of a docker container
		//tURL, err := url.Parse("http://"+"localhost:"+ se)
		tURL, err := url.Parse("http://" + se)
		
		if err != nil{
			logCh <- "Balancer: Cannot Parse URL:" + se + err.Error()
			return
		}

		req.Host = tURL.Host
		req.URL.Host = tURL.Host
		req.URL.Scheme = tURL.Scheme
		req.RequestURI = ""

		s, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil{
			logCh <- "Balancer: Cannot SplitHostPort " + req.RemoteAddr
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil{
			rw.WriteHeader(http.StatusInternalServerError)
			logCh <- "Balancer: Cannot process request: " + err.Error()
			return
		}

		for key, values := range resp.Header{
			for _,  value := range values{
				rw.Header().Set(key, value)
			}
		}
		rw.WriteHeader(resp.StatusCode)
		io.Copy(rw, resp.Body)
		logCh <- "Balancer: Forwarded to: " + s
	})

	fmt.Printf("Balancer listening from %s\n", config.Port)
	http.ListenAndServe(":" + config.Port, proxy)
}
func handleLog(){
	for{
		//conn, err := net.Dial("tcp", "127.0.0.1:" + config.LoggerPort)
		conn, err := net.Dial("tcp", "172.17.0.1:" + config.LoggerPort)
		if err != nil{
			log.Fatalf("Logger not loaded %s\n", err)
		}
		conn.Write([]byte(<-logCh))
	}
	
}