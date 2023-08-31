package main

import "fmt"

func main() {
	// declare channel and set capacity of buffered channel
	ch := make(chan string, 10)

	// giving data
	ch <- "First message"
	ch <- "Second message"

	// printing data
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
