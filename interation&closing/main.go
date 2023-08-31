package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// declare channel
	ch := make(chan string)
	// run function with goroutine
	go ThrowStars(ch)

	// print out with for loop correctly
	for {
		// init message string & open boolean
		message, open := <-ch
		// check if its still open
		if !open {
			break
		}
		// print
		fmt.Println(message)
	}
}

// function itself
func ThrowStars(ch chan string) {
	// number of rounds that we need to throw stars
	numberOfRounds := 3
	// loop of as many number of rounds
	for i := 0; i < numberOfRounds; i++ {
		// random number
		score := rand.Intn(10)
		// sending string with result through the channel
		ch <- fmt.Sprint("You scored:", score)
	}
	// close channel so there is wouldn't be deadlock
	close(ch)
}

// func main() {
// 	#declare variable and their types
// 	ch := make(chan string)
// 	numberOfRounds := 3

// 	#run function with goroutine
// 	go ThrowStars(ch, numberOfRounds)

// 	#print out as many number of rounds
// 	for i := 0; i < numberOfRounds; i++ {
// 		fmt.Println(<-ch)
// 	}
// }

// #function that receive channel
// func ThrowStars(ch chan string, numberOfRounds int) {
// 	#throw stars as many number of rounds
// 	for i := 0; i < numberOfRounds; i++ {
// 		#random numbers
// 		score := rand.Intn(10)
// 		#sending string with result through the channel
// 		ch <- fmt.Sprint("You scored:", score)
// 	}
// }
