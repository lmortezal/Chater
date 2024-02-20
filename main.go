package main

import (
	"github.com/lmortezal/ChatOnline/cmd"
	s "github.com/lmortezal/ChatOnline/server"
	"github.com/lmortezal/ChatOnline/client"
)

func main() {
	domain, port, server := cmd.Execute()
	if server {
		s.Startlistening(domain, port)
	} else if !server{
		client.Startconnection(domain, port)
	}
}