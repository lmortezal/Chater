package client

import (
	"bufio"
	"fmt"
	//"os"
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


func Startconnection(domain string , port int, messages <-chan  MsgStruct , serverMsg chan MsgStruct){

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
	Name := <-messages
	fmt.Fprint(conn_tls, Name.Name + "\n")
	
	go func() {
		input := bufio.NewScanner(conn_tls)
		for input.Scan(){
			serverMsg <- MsgStruct{Message: strings.Split(input.Text(), ":")[1], Name: strings.Split(input.Text(), ":")[0]}}
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
	
