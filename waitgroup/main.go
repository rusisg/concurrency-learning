package main

import (
	"fmt"
	"sync"
)

func main() {
	// calling waitgroup from sync package
	var beeper sync.WaitGroup
	evilNinjas := []string{"Rasul", "Niger", "Asylbek"}
	// adding how many goroutines will wait (len of evilNinjas)
	beeper.Add(len(evilNinjas))
	// by running function with for loop attacking each of evilNinjas
	for _, evilNinja := range evilNinjas {
		go attack(evilNinja, &beeper)
	}
	beeper.Wait()

	// this print to show that there is another goroutine which is mainGoroutine
	fmt.Println("Mission completed")
}

// function itself and also converting beeper back to data
func attack(evilNinja string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked evil ninja:", evilNinja)
	// after finishing function's logic waitgroup would be done
	beeper.Done()
}

// func main() {
// 	// calling waitgroup from sync package
// 	var beeper sync.WaitGroup

// 	// adding how many goroutines will wait (1)
// 	beeper.Add(1)
// 	// run function with goroutine and converting beeper into pointer
// 	go attack("Rasul", &beeper)
// 	beeper.Wait()

// 	// this print to show that there is another goroutine which is mainGoroutine
// 	fmt.Println("Mission completed")
// }

// // function itself and also converting beeper back to data
// func attack(evilNinja string, beeper *sync.WaitGroup) {
// 	fmt.Println("Attacked evil ninja:", evilNinja)
// 	// after finishing function's logic waitgroup would be done
// 	beeper.Done()
// }
