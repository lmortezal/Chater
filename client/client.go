package client

import (
	"fmt"
	"net"
)

var (
	msg string  
	Name string
)

func Startconnection(domain string , port int){

	// connect to this socket
	var addr = fmt.Sprintf("%s:%d", domain, port)
	// for i := 1; i < 5 ; i++{
	// 	//time.Sleep(time.Duration(i * 10) * time.Second)
	// 	for j := i * 10; j > 0 ; j--{
	// 		fmt.Printf("Trying to connect to %v in %d Seconds\r",  addr ,j)
	// 		time.Sleep(time.Second) // Adjust the duration as needed
	// 	}
	// 	fmt.Println()
	// 	fmt.Print("\033[H\033[2J")
	// }
		
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	fmt.Printf("connecting to %v\nEnter your name:\n" , addr)
	fmt.Scan(&Name)
	fmt.Fprintf(conn, Name + "\n")
	
	go printlisten(conn)


	for {
		if err != nil {
			// handle error
			fmt.Println("Error:", err.Error())
		}
		fmt.Scan(&msg)
		fmt.Fprintf(conn, msg)
	}
	
	
	
	
}
	
func printlisten(conn net.Conn){
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error:", err.Error())
			return
		}
		fmt.Println(string(buf))
	}
}