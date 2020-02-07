package config

import (
)

//StartCommands is a slice of strings that refer to all the commands needed to launch the server
var StartCommands []string
//LoggerPort holds the port value that points to the logger program
var LoggerPort string = "9090"

func init(){

}
//GetCommands returns the StartCommands slice
func GetCommands()[]string{
	return StartCommands
}