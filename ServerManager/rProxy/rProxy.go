package rproxy

import (
	"fmt"
	"os/exec"
)

//StartProxy starts a reverse proxy pointing to the load balancer
func StartProxy(logCh chan string){
	cmd := exec.Command("./Proxy/proxyServer")
	//cmd.Start()
	
	fmt.Println("Proxy started")
	// err := cmd.Wait()
	str, err := cmd.CombinedOutput()
	if err != nil{
		logCh <- "Manager: proxy failed" + err.Error()
	} else {
		logCh <- "Manager: " + string(str) 
	}
}