package main

import (
	"fmt"
	"sync"
	"time"
)

type Chopstick struct{ sync.Mutex }
type Philosopher struct {
	leftstick, rightstick *Chopstick
}

func (p Philosopher) eat(id int, sem chan int) {
	p.leftstick.Lock()
	p.rightstick.Lock()
	fmt.Printf("Starting to eat %v\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Finishing eating %v\n", id)
	p.leftstick.Unlock()
	p.rightstick.Unlock()

	// Remove item from channel
	<-sem
}

func dininghost(id int, p Philosopher) {

}

func main() {
	// Implement the dining philosopher’s problem with the following constraints/modifications.
	// 1. There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
	// 2. Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
	// 3. The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
	// 4. In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
	// 5. The host allows no more than 2 philosophers to eat concurrently.
	// 6. Each philosopher is numbered, 1 through 5.
	// 7. When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself,
	// where <number> is the number of the philosopher.
	// 8. When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself,
	// where <number> is the number of the philosopher.
	var totalPeople int = 5
	const workerThreads int = 2

	sem := make(chan int, workerThreads)

	sticks := make([]*Chopstick, totalPeople)
	philosophers := make([]*Philosopher, totalPeople)

	// init chopsticks
	for i := 0; i < 5; i++ {
		sticks[i] = new(Chopstick)
	}

	// Init philosophers
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{sticks[i], sticks[(i+1)%5]}
	}

	// Process meal
	for i := 0; i < 3; i++ {
		// fmt.Printf("Times to eat: %v\n", i)
		for ii := 0; ii < 5; ii++ {
			// Push 1 on channel; should block if channel is at 2
			sem <- 1
			// fmt.Printf("Philosopher: %v\n", ii+1)
			go philosophers[ii].eat(ii+1, sem)

		}
		time.Sleep(time.Second)
	}
}
