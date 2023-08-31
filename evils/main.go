package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	evils := []string{"Madi", "Rasul", "Arsen", "Nurdaulet", "Asylbek", "Mama"}

	for _, evil := range evils {
		go attack(evil)
	}

	time.Sleep(time.Second * 1)
}

func attack(target string) {
	fmt.Printf("Throwing stars at %v\n", target)
	time.Sleep(time.Second)
}
