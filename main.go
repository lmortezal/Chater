package main

import (
	"github.com/lmortezal/Chater/cmd"
	s "github.com/lmortezal/Chater/server"
	"github.com/lmortezal/Chater/client"
)

func main() {
	domain, port, server := cmd.Execute()
	if server {
		s.Startlistening(domain, port)
	} else if !server{
		client.Startconnection(domain, port)
	} else{
		panic("Error")
	}
}