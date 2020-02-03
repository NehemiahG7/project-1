package config

import (
	"flag"
)

//FirewallPort holds the value for the fP flag
var FirewallPort string

//LoadPort holds the value for the lP flag
var LoadPort string

//LoggerPort holds the value for the lP flag
var LoggerPort string

func init(){
	flag.StringVar(&FirewallPort, "fP", "6060", "Use fP to designate the port that the firewall listens on")
	flag.StringVar(&LoadPort, "lP", "7070", "Use lP to designate the port that the firewall sends connections to")
	flag.StringVar(&LoggerPort, "lP", "9090", "Use lP to designate the port that the proxy looks for a logger on")


	flag.Parse()
}