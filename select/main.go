package main

import "fmt"

func main() {
	// declare two channels and their types
	ninja1, ninja2 := make(chan string), make(chan string)

	// run function with goroutine
	go captainElect(ninja1, "KOTAK1")
	go captainElect(ninja2, "KOTAK2")

	// select is chooses one of the channels that will be running and shows result
	select {
	// passing data of ninja1 into message1 variable
	case message1 := <-ninja1:
		fmt.Println(message1)
	// passing data of ninja1 into message2 variable
	case message2 := <-ninja2:
		fmt.Println(message2)
	}

	roughlyFair()
}

// function that will put message into ninja variable
func captainElect(ninja chan string, message string) {
	ninja <- message
}

func roughlyFair() {
	// initialize two channels with no matter types and they will close immediately
	ninja1 := make(chan interface{})
	close(ninja1)
	ninja2 := make(chan interface{})
	close(ninja2)
	// there is inits
	var ninja1Count, ninja2Count int

	// for loop that will runs 1000 times
	for i := 0; i < 1000; i++ {
		// select is chooses one of the channels that will be running and shows result
		select {
		// this channel or...
		case <-ninja1:
			ninja1Count++
		// ...or this channel
		case <-ninja2:
			ninja2Count++
		}
	}

	// printing
	fmt.Printf("ninja1Count: %d ninja2Count: %d\n", ninja1Count, ninja2Count)
}
