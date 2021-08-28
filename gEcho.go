package main

import (
	"log"
	"net"
)

//echo func echoes received data
func echo(conn net.Conn) {
	defer conn.Close()

	//buffer to store received data
	b := make([]byte, 512) 
	for {
		//receive data via conn.Read into a buffer
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("Client Disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected Error")
			break
		}
		log.Printf("Received %d bytes: %s\n", size, string(b))

		//send data via conn.Write
		log.Println("Writing Data")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Fatalln("Unable to Write Data")
		}
	}
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