package main

import (
	"github.com/lmortezal/ChatOnline/cmd"
	s "github.com/lmortezal/ChatOnline/server"
)

func main() {
	domain, port, server := cmd.Execute()
	if server {
		s.Startlistening(domain, port)
	}
}