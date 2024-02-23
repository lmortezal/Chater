package main

import (
	"time"

	tui "github.com/lmortezal/Chater/Tui"
	"github.com/lmortezal/Chater/client"
	"github.com/lmortezal/Chater/cmd"
	s "github.com/lmortezal/Chater/server"
)

func main() {
	tui.Tui_main()
	time.Sleep(1000 * time.Second)
	domain, port, server := cmd.Execute()
	if server {
		s.Startlistening(domain, port)
	} else if !server{
		client.Startconnection(domain, port)
	} else{
		panic("Error")
	}
}