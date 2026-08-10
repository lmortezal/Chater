package server

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"sync"
)

var (
	messages = make(chan string)
	clients  = make(map[Client]bool)
	// messages2 = make(chan allmessages)
	mu sync.RWMutex
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

func Startlistening(domain string, port int) error {
	if domain == "" {
		domain = "127.0.0.1"
	}
	if port == 0 {
		port = 8080
	}
	cert, err := tls.LoadX509KeyPair("./server/server.crt", "./server/server.key")
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	config_tls := tls.Config{Certificates: []tls.Certificate{cert}}

	var addr = fmt.Sprintf("%s:%d", domain, port)
	l_tld, err := tls.Listen("tcp4", addr, &config_tls)

	if err != nil {
		log.Fatal(err)
		return err
	} else {
		fmt.Println("Listening on ", addr)
	}

	// Close the listener when the application closes. and handle the error
	// each client send to server give it and print in the server and return to the client
	defer l_tld.Close()

	go Boradcast()

	for {
		// Listen for an incoming connection.
		//conn, err := l.Accept()
		conn_tld, err := l_tld.Accept()
		if err != nil {
			log.Printf("Accept error: %v", err)
			continue
		}

		// Handle connections in a new goroutine.
		go handleRequest(conn_tld)
	}
}

// just Boradcast the message to all clients
func Boradcast() {
	for msg := range messages {
		mu.RLock()
		for cli := range clients {
			fmt.Fprint(cli.Conn, msg+"\n")
		}
		mu.RUnlock()
	}
	defer func() {
		fmt.Println("Boradcast is closed")
	}()
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	input := bufio.NewScanner(conn)
	fmt.Println("New connection from:", conn.RemoteAddr())
	var name string
	for input.Scan() {
		name = input.Text()
		messages <- name + " has joined"
		break
	}

	cli := Client{conn, name}
	mu.Lock()
	clients[cli] = true
	mu.Unlock()

	defer func() {
		mu.Lock()
		delete(clients, cli)
		mu.Unlock()
		messages <- name + " has left"
		conn.Close()

	}()

	for input.Scan() {
		if input.Text() == "" {
			continue
		}
		messages <- name + ":" + input.Text()
		fmt.Println("Received from ", name, ":", input.Text())
	}
}
