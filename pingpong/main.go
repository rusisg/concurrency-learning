package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "PONG%s", time.Now())
}

func main() {
	http.HandleFunc("/ping", pong)

	log.Println("Start server!")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Printf("Error while starting server: %e", err)
	}
}
