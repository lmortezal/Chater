package client

import (
	"bufio"
	"fmt"
	"crypto/tls"
	"strings"
)

var (
	msg string  
	Name string
)

type MsgStruct struct {
    Name    string
    Message string
}


func Startconnection(domain string , port int, messages <-chan  MsgStruct , serverMsg chan<- MsgStruct){

	// connect to this socket
	var addr = fmt.Sprintf("%s:%d", domain, port)
	//TODO // add check connection to the server 


	conn_tls, err := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true} )
	if err != nil {
		fmt.Println("Error:", err.Error())
		//fmt.Println("Line 31")
		return
	}
	//fmt.Printf("connecting to %v\nEnter your name:\n" , addr)
	
	// enter name from chan messages
	get_tui := <-messages
	fmt.Fprint(conn_tls, get_tui.Name + "\n")
	
	go func() {
		input := bufio.NewScanner(conn_tls)
		for input.Scan(){
			if strings.Split(input.Text(), ":")[0] == get_tui.Name  {
				continue
			}
			parts := strings.Split(input.Text(), ":")
			if len(parts) >= 2 {
				serverMsg <- MsgStruct{Message: parts[1], Name: parts[0]}
			} else if len(parts) == 1 {
				serverMsg <- MsgStruct{Message: "", Name: parts[0]}
			} else {
				// handle the case where parts is empty, if necessary
				fmt.Println("Error: parts is empty")
			}
		}
	}()

	
	//fmt.Scan(&Name)
	
	
	// get message from the Tui(message channel) and send it to the server
	for {
		select {
		case msg := <-messages:
			fmt.Fprintf(conn_tls, msg.Message + "\n")
		}
	
	}

	// terminal := bufio.NewScanner(os.Stdin)
	// //fmt.Println("Enter your message:")
	// for terminal.Scan(){
	// 	msg = terminal.Text()
	// 	fmt.Fprintf(conn_tls, msg + "\n")
		
	// }
	
}
	
