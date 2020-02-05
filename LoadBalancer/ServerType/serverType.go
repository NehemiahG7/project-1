package serverhand

import (
	"strings"
)

//Server is a struct representing a server
type Server struct{
	//Name holds the name of the docker servers running on this network
	Name string
	//TotalCon represents the total connections made to the server since the server launched
	TotalCon int
}

//LoadServers takes in a string of names seperated by a coma and makes each name into a server struct, returning the slice
func LoadServers(servs string) *[]Server{
	arry := strings.Split(servs, ",")
	servers := make([]Server, 0)
	for i :=0; i < len(arry); i++{
		serv := Server{arry[i],0}
		servers = append(servers, serv)
	}
	return &servers
}
//GetServer checks the currently active servers and returns the one with the least use
func GetServer(servs []Server) string{
	server := servs[0]
	for i := 1; i < len(servs); i++{
		if server.TotalCon > servs[i].TotalCon{
			server = servs[i]
		}
	}
	return server.Name
}