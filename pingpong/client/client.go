package client

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080/ping")
	if err != nil {
		log.Println("Error connecting to server")
		return
	}
	defer conn.Close()
}
