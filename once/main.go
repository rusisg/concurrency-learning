package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var missionCompleted bool

func main() {
	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			if foundTreasure() {
				markMissionCompleted()
			}
			wg.Done()
		}()
	}
	wg.Wait()

	checkMissionCompletion()
}

func checkMissionCompletion() {
	if missionCompleted {
		fmt.Println("Mission is completed")
	} else {
		fmt.Println("Mission is failure")
	}
}

func markMissionCompleted() {
	missionCompleted = true
	fmt.Println("Mission is now completed")
}

func foundTreasure() bool {
	return 0 == rand.Intn(10)
}

// var missionCompleted bool
// var numberOfGoroutines int

// func main() {
// 	var wg sync.WaitGroup
// 	numberOfGoroutines = 10
// 	wg.Add(numberOfGoroutines)

// 	for i := 0; i < numberOfGoroutines; i++ {
// 		go func() {
// 			if foundTreasure() {
// 				markMissionCompleted()
// 			}
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()

// 	checkMissionCompletion()
// }

// func checkMissionCompletion() {
// 	if missionCompleted {
// 		fmt.Println("Mission is completed")
// 	} else {
// 		fmt.Println("Mission is failure")
// 	}
// }

// func markMissionCompleted() {
// 	missionCompleted = true
// 	fmt.Println("Mission is now completed")
// }

// func foundTreasure() bool {
// 	return 0 == rand.Intn(10)
// }
