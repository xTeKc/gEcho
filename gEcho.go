package main

import (
	"log"
	"net"
)

//echo func echoes received data
func echo(conn net.Conn) {
	defer conn.Close()
}