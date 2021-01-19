package main 

import ( 
	"fmt"
	"time"
) 

func add(value int) { 
	for i:=0; i<50; i++ { 
		fmt.Println(value+i);
	}
} 

// There are 2 goroutines, calling a function to print numbers from 1-50
// and 1000-1049. But we do not know which goroutine will be scheduled to 
// start first, or when a goroutine would end.
// Running this n times, might give us n different output
// There is no guarantee when each of the gorounites become active
func main() {
	// Initiate first goroutine. There is no guarantee that this 
	// is the first goroutine to start and end.
	go add(1) 
	
	// The second goroutine which might start executing at any 
	// point, ["before", "during" or "after"] the first goroutine.
	go add(1000)
	
	// Adding a sleep to make sure that the two routines end and 
	// values are printed
	time.Sleep(10 * time.Second)
} 
