package client

import (
	"net"
	"fmt"
)


func Startconnection(ip string , domain int){
	fmt.Println("1")

	conn , err := net.Dial("tcp4" , "localhost:8081")

	data := []byte("Hello, Server!")
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	 

}
	