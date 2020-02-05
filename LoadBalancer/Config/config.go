package config

import (
	"flag"
	"github.com/NehemiahG7/project-1/LoadBalancer/ServerType"

)

//LoggerPort holds the value for the lP flag
var LoggerPort string

//Names holds the value for the addr flag
var Names string

//Servers is a pointer to a [] of Servers containing all servers this balancer goes between
var Servers *[]serverhand.Server

func init(){
	flag.StringVar(&LoggerPort, "lP", "9090", "Use lP to designate the port that the logger listens on")
	flag.StringVar(&Names, "addr", "8080", "Use addr to give a seriese of ports for the logger to balance between. Seperate each port by only a coma")
	flag.Parse()

	Servers = serverhand.LoadServers(Names)
}