package main

import (
	tui "github.com/lmortezal/Chater/Tui"
	"github.com/lmortezal/Chater/client"
	"github.com/lmortezal/Chater/cmd"
	s "github.com/lmortezal/Chater/server"
)

func main() {
	tui.Tui_main()
	domain, port, server := cmd.Execute()
	if server {
		s.Startlistening(domain, port)
	} else if !server{
		client.Startconnection(domain, port , nil , nil)
	} else{
		panic("Error")
	}
}