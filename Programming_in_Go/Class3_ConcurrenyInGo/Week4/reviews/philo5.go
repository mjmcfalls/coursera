package main

import (
	"fmt"
	"sync"
)

const (
	guests   = 5 // number of participating philosophers
	servings = 3 // allowed number of servings per guest
	capacity = 2 // max number of guest allowed to eat at the same time
)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	id                            int
	leftChopstick, rightChopstick *chopstick
}

func (p philosopher) eat(pWg *sync.WaitGroup) {
	defer pWg.Done()

	p.leftChopstick.Lock()
	p.rightChopstick.Lock()

	fmt.Printf("starting to eat %v\n", p.id)
	fmt.Printf("finishing eating %v\n", p.id)

	p.rightChopstick.Unlock()
	p.leftChopstick.Unlock()

}

func xDinningPhilosophers() {
	chopsticks := make([]*chopstick, guests)
	for j := 0; j < guests; j++ {
		chopsticks[j] = new(chopstick)
	}

	philosophers := make([]*philosopher, guests)
	for j := 0; j < guests; j++ {
		philosophers[j] = &philosopher{j + 1, chopsticks[j], chopsticks[(j+1)%guests]}
	}

	var wg sync.WaitGroup
	wg.Add(guests)
	for j := 0; j < guests; j++ {
		go philosophers[j].eat(&wg)
	}
	wg.Wait()
}

var (
	semaphoreSize = capacity

	mu                 sync.Mutex
	totalTasks         int
	curConcurrentTasks int
	maxConcurrentTasks int
)

func timeConsumingTask() {
	mu.Lock()
	totalTasks++
	curConcurrentTasks++
	if curConcurrentTasks > maxConcurrentTasks {
		maxConcurrentTasks = curConcurrentTasks
	}
	mu.Unlock()

	// the time consuming task
	xDinningPhilosophers()

	mu.Lock()
	curConcurrentTasks--
	mu.Unlock()
}

func main() {
	sem := make(chan struct{}, semaphoreSize)
	var wg sync.WaitGroup
	for i := 0; i < servings; i++ {
		// acquire semaphore
		sem <- struct{}{}
		wg.Add(1)

		// the host goroutine
		go func() {
			timeConsumingTask()
			// release semaphore
			<-sem
			wg.Done()
		}()
	}

	// wait for all tasks to finish
	wg.Wait()
}
