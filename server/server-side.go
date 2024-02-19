package server

import (
	"fmt"
	"log"
	"net"
	"strings"
)


func Startlistening(domain string, port int) {
	if domain == "" {
		domain = "127.0.0.1"
	}
	if port == 0 {
		port = 8080
	}
	
	var addr = domain + ":" + fmt.Sprintf("%d", port)
	fmt.Println("Listening on ", addr)
	l, err := net.Listen("tcp4", addr)
	if err != nil {
		log.Fatal(err)
	}
	// Close the listener when the application closes. and handle the error 
	// each client send to server give it and print in the server and return to the client
	defer l.Close()
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}


}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// print data in conn ( data is a buffer) 
	// print human readable data
	fmt.Println("New connection from:", conn.RemoteAddr())
	


	conn.Write([]byte("Connection is success \nYour Ip Address is: " + strings.Split(conn.RemoteAddr().String(), ":")[0] + "\n"))
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	for {
		// Read the incoming connection into the buffer.
		reqLen, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading:", err.Error())
			break
		}
		// Convert the buffer to a string and print it.
		reqStr := string(buf[:reqLen])
		fmt.Println("Received from client:", reqStr)
		// Send the received string back to the client.
		conn.Write([]byte("Message received: " + reqStr))
	}
}