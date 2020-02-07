package main

import(
	"github.com/NehemiahG7/project-1/ServerManager/Config"
	"os/exec"
	"log"
	"fmt"
	"net"
	"strings"
)

//Channel to pass messages for the logger
var logCh chan string = make(chan string)

var cmdCh chan string = make(chan string)

var power chan string = make(chan string)


func main(){
	go startLogger()
	fmt.Println(<- cmdCh)
	go handleLog()
	logCh <- "Manager Online"
	go startProxy()
	go startBalancer()
	go startServer()
	
	<- power
}
func startLogger(){
	cmd := exec.Command("./Logger/logger")
	cmd.Start()
	
	cmdCh <- "Logger started"
	
	err := cmd.Wait()
	if err != nil{
		log.Fatalf("Logger could not be started, closing program: %s\n", err)
	}
}
func startProxy(){
	cmd := exec.Command("./Proxy/proxyServer")
	//cmd.Start()
	
	fmt.Println("Proxy started")
	// err := cmd.Wait()
	str, err := cmd.CombinedOutput()
	if err != nil{
		logCh <- "Proxy could not be started, closing program: %s\n" + err.Error()
	} else {
		logCh <- "Manager: " + string(str) 
	}
}
func startBalancer(){
	//Split command into command and args
	s := "docker run --rm -p 8082:9090 --name balancer --network my-net balancer -addr=tac1,tac0"
	args := strings.Split(s, " ")

	//Create command with args[0] as command and rest as args
	cmd := exec.Command(args[0], args[1:]...)
	// cmd.Start()
	
	fmt.Println("Balancer started")
	// err := cmd.Wait()
	str, err := cmd.CombinedOutput()
	if err != nil{
		logCh <- "Proxy could not be started, closing program: %s\n" + err.Error()
	} else {
		logCh <- "Manager: " + string(str) 
	}
}
func startServer(){
	//Split command into command and args
	s := "docker run --rm -d --name tac0 --network my-net tictactoe"
	args := strings.Split(s, " ")
	
	//Create command with args[0] as command and rest as args
	cmd := exec.Command(args[0], args[1:]...)
	// cmd.Start()
	
	fmt.Println("Server started")
	// err := cmd.Wait()
	str, err := cmd.CombinedOutput()
	if err != nil{
		logCh <- "Proxy could not be started, closing program: %s\n" + err.Error()
	} else {
		logCh <- "Manager: " + string(str) 
	}
}
func handleLog(){
	for{
		//conn, err := net.Dial("tcp", "127.0.0.1:" + config.LoggerPort)
		conn, err := net.Dial("tcp", "localhost:" + config.LoggerPort)
		if err != nil{
			log.Fatalf("Logger not loaded %s\n", err)
		}
		conn.Write([]byte(<-logCh))
	}
	
}
