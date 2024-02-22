package client

import (
	"bufio"
	"fmt"
	"os"
	"crypto/tls"
)

var (
	msg string  
	Name string
)

func Startconnection(domain string , port int){

	// connect to this socket
	var addr = fmt.Sprintf("%s:%d", domain, port)
	//TODO // add check connection to the server 


	conn_tls, err2 := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true} )
	if err2 != nil {
		fmt.Println("Error:", err2.Error())
		return
	}
	fmt.Printf("connecting to %v\nEnter your name:\n" , addr)
	go func() {
		input := bufio.NewScanner(conn_tls)
		for input.Scan(){
			fmt.Println(input.Text())
		}
	}()

	
	fmt.Scan(&Name)
	fmt.Fprintf(conn_tls, Name)
	

	terminal := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your message:")
	for terminal.Scan(){
		msg = terminal.Text()
		fmt.Fprintf(conn_tls, msg + "\n")
	}
	
}
	
