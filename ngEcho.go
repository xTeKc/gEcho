package main

import (
	"log"
	"net"
)

//echo func echoes received data
func echo(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln("Unable to Read Data")
	}
	log.Printf("Read %d bytes: %s", len(s), s)

	log.Println("Writing Data")
	writer := bufio.NewWriter(conn)
	if _, err := writer.WriterString(s); err != nil {
		log.Fatalln("Unable to Write Data")
	}
	writer.Flush()
}


//alt echo func using Copy()
// func echo (conn, net.Conn) {
// 	defer conn.Close()
// 	//copy data from io.Reader to io.Writer via io.Copy()
// 	if _, err := Copy(conn, conn); err != nil {
// 		log.Fatalln("Unable to Read/Write Data")
// 	}
// }


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