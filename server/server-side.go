package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"crypto/tls"
)

var (
	messages = make(chan string)
	clients = make(map[Client]bool)
	// messages2 = make(chan allmessages)
)

// TODO  // create a struct for the messages to filter dublicate message
// type allmessages struct {
// 	Conn net.Conn
// 	Message string
// }

type Client struct {
	Conn net.Conn
	Name string
}



func Startlistening(domain string, port int) {
	if domain == "" {
		domain = "127.0.0.1"
	}
	if port == 0 {
		port = 8080
	}

	
	var addr = fmt.Sprintf("%s:%d", domain, port)
	l, err := net.Listen("tcp4", addr)
	l_tld , err2 := tls.Listen("tcp4", addr, )

	if err2 != nil {
		log.Fatal(err2)
	} else{
		fmt.Println("Listening on ", addr)
	}

	if err != nil {
		log.Fatal(err)
	} else{
		fmt.Println("Listening on ", addr)
	}
	// Close the listener when the application closes. and handle the error 
	// each client send to server give it and print in the server and return to the client
	defer l.Close()
	defer l_tld.Close()



	go Boradcast()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		conn_tld, err2 := l_tld.Accept()
		if (err2 != nil) && (err != nil){
			log.Fatal(err2)
		}

		// Handle connections in a new goroutine.
		// clients[Client{conn, ""}] = true // this is a bug
		if conn != nil {
			go handleRequest(conn)
		}else {
			go handleRequestTls(conn_tld)
		}
		

	}


}


// just Boradcast the message to all clients
func Boradcast(){
		for msg := range messages {
			for cli := range clients {
				fmt.Fprint(cli.Conn , msg + "\n")
			}
		}
}
func handleRequestTls(conn net.Conn){
	input := bufio.NewScanner(conn)
	fmt.Println("New connection from:", conn.RemoteAddr())
	var name string
	for input.Scan(){
		name = input.Text()
		break
	}
	
	cli := Client{conn, name}
	clients[cli] = true

	for input.Scan(){
		if input.Text() == ""{
			continue
		}
		messages <- name + " say : " + input.Text()
		fmt.Println("Received from ", name, ":", input.Text())
	}
	

}
// Handles incoming requests.
func handleRequest(conn net.Conn) {
	input := bufio.NewScanner(conn)
	fmt.Println("New connection from:", conn.RemoteAddr())
	var name string
	for input.Scan(){
		name = input.Text()
		break
	}
	
	cli := Client{conn, name}
	clients[cli] = true



	for input.Scan(){
		if input.Text() == ""{
			continue
		}
		messages <- name + " say : " + input.Text()
		fmt.Println("Received from ", name, ":", input.Text())
	}
}