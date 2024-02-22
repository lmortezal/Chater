package client

import (
	"bufio"
	"fmt"
	// "net"
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
	// for i := 1; i < 5 ; i++{
	// 	//time.Sleep(time.Duration(i * 10) * time.Second)
	// 	for j := i * 10; j > 0 ; j--{
	// 		fmt.Printf("Trying to connect to %v in %d Seconds\r",  addr ,j)
	// 		time.Sleep(time.Second) // Adjust the duration as needed
	// 	}
	// 	fmt.Println()
	// 	fmt.Print("\033[H\033[2J")
	// }
	conn_tls, err2 := tls.Dial("tcp", addr, &tls.Config{InsecureSkipVerify: true} )
	//conn, err := net.Dial("tcp", addr)

	if err2 != nil {
		fmt.Println("Error:", err2.Error())
		return
	}
	fmt.Printf("connecting to %v\nEnter your name:\n" , addr)
	fmt.Scan(&Name)
	fmt.Fprintf(conn_tls, Name)
	
	go func() {
		input := bufio.NewScanner(conn_tls)
		for input.Scan(){
			fmt.Println(input.Text())
		}
	}()

	terminal := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your message:")
	for terminal.Scan(){
		msg = terminal.Text()
		fmt.Fprintf(conn_tls, msg + "\n")
	}
	
}
	
