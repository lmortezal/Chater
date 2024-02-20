
package client

import (
	"net"
	"fmt"
)

func Startconnection(domain string , port int){
	// connect to this socket
	var addr = fmt.Sprintf("%s:%d", domain, port)
	var msg string  
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}

	fmt.Printf("connecting to %v :\nNow you can type it !\n" , addr)
	for {
		if err != nil {
			// handle error
			fmt.Println("Error:", err.Error())
		}
		fmt.Scan(&msg)
		fmt.Fprintf(conn, msg + "\n")
	}
	

}
	