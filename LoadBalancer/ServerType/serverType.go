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
func LoadServers(servs string) *map[string]int{
	arry := strings.Split(servs, ",")

	servers := make(map[string]int)
	for i :=0; i < len(arry); i++{
		servers[arry[i]] = 0
	}
	return &servers
}
//GetServer checks the currently active servers and returns the one with the least use
func GetServer(servs map[string]int) string{

	j := -1
	i := 0
	var str string
	for k := range servs{
		if servs[k] <= i || j < 0{
			j++
			str = k
			i = servs[k]
		}
	}
	return str

}

//GetKeys returns all keys in mymap
func GetKeys(mymap map[string]int) string {
    keys := ""
    for k := range mymap {
        keys = keys + "," + k
	}
	return keys
}