package main

import (
	"fmt" 
	tui "github.com/lmortezal/Chater/Tui"
	"github.com/lmortezal/Chater/cmd"
	s "github.com/lmortezal/Chater/server"
)

func main() {
	domain, port, server := cmd.Execute()
	if server {
		s.Startlistening(domain, port)
	} else if !server{
		tui.Tui_main(domain , port)
	} else{
		fmt.Println("Error: server or client mode not set.")
	}
}