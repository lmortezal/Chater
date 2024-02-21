package server

import (
	"fmt"
	"log"
	"net"
)

var (
	messages = make(chan string)
	clients = make(map[Client]bool)
)

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
	

	if err != nil {
		log.Fatal(err)
	} else{
		fmt.Println("Listening on ", addr)
	}
	// Close the listener when the application closes. and handle the error 
	// each client send to server give it and print in the server and return to the client
	defer l.Close()


	go Boradcast()

	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle connections in a new goroutine.
		clients[Client{conn, ""}] = true
		go handleRequest(conn)

	}


}

func Boradcast(){
		for msg := range messages {
			for cli := range clients {
				fmt.Fprint(cli.Conn , msg)
			}
		}
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// print data in conn ( data is a buffer) 
	// print human readable data

	fmt.Println("New connection from:", conn.RemoteAddr())
	
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)
	var name string
	
	for {
		// Read the incoming connection into the buffer.
		if clients[Client{conn, ""}]{
			fmt.Println("New client")
			_, err := conn.Read(buf)
			if err != nil {
				log.Println("Error reading:", err.Error())
				break
			}
			
			name = string(buf[:])
			fmt.Printf("Name received: %s", name)
			for cli := range clients {
				if cli.Conn == conn{
					cli.Name = name
				}
			}
			clients[Client{conn, ""}] = false
			continue
		}


		reqLen, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading:", err.Error())
			break
		}
		// Convert the buffer to a string and print it.
		reqStr := string(buf[:reqLen])
		messages <- reqStr
		fmt.Println(&messages)
		fmt.Println("Received from ", name, ":", reqStr)
		// Send the received string back to the client.
		//conn.Write([]byte("Message received: " + reqStr))
	}
}