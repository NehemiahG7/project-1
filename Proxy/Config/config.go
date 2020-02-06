package config

import (
	"flag"
)

//ProxyPort holds the value for the pP flag
var ProxyPort string

//LoggerPort holds the value for the lP flag
var LoggerPort string

//BalancerPort holds the value for the sP flag
var BalancerPort string

func init(){
	flag.StringVar(&ProxyPort, "pP", "8081", "Use pP to designate the port that the proxy listens on")
	flag.StringVar(&LoggerPort, "lP", "9090", "Use lP to designate the port that the proxy looks for a logger on")
	flag.StringVar(&BalancerPort, "sP", "8082", "Use sP to designate the port that the proxy looks for a server on")
	flag.Parse()


}