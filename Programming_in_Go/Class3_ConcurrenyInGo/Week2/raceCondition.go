package main

import (
	"fmt"
	"time"
)

func createRace(c *int) {
	*c++
}
func readRace(c *int) {
	fmt.Printf("%v", *c)
}

var count int

func main() {
	// Race condition caused by reading and writing to the same memory location at the same time
	go createRace(&count)
	go readRace(&count)
	time.Sleep(1 * time.Second)
}
