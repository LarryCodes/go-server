package main

import (
	"connectionHandler"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Could not create listener:", err)
	}
	defer listener.Close()

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("Could not listen to connection at this time:", err)
		}
		go connectionHandler.Handle(connection)
	}
}
