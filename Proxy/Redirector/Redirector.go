package balancer

import (
	"github.com/NehemiahG7/project-1/Proxy/Config"
)

//GetPort returns the port of a server
func GetPort()string {
	return config.BalancerPort
}