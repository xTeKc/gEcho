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