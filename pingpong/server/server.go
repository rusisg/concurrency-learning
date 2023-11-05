package server

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080/ping")
	if err != nil {
		log.Println("Error connecting while listening client")
	}

	defer listener.Close()

	fmt.Println("Server is listening on localhost:8080/ping")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection")
			continue
		}
		go HandleConnection(conn)
	}

}

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("Start server!")

	_, err := conn.Write([]byte("ping\n"))
	if err != nil {
		log.Println("Error sending a data")
	}

}
