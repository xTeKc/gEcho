package main

import (
	"log"
	"net"
)

//echo func echoes received data
func echo(conn net.Conn) {
	defer conn.Close()
}


func main() {
	//bind to TCP port 20080 on all interfaces
	listener, err := net.Listen("tcp", ":20080")
	if err != nil {
		log.Fatalln("Unable to Bind to Port")
	}
	log.Println("Listening on 0.0.0.0:20080")
	for {
		//wait for connection and create net.Conn on connection
		conn, err := listener.Accept()
		log.Println("Received Connection")
		if err != nil {
			log.Fatalln("Unable to Accept Connection")
		}
		//handle the connection while using go-routine for concurrency
		go echo(conn)
	}
}