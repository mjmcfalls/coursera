/**
The following code will call a function addX, which has a number of concurrent assignments for x.
It will return x + 1, which, based on which anonymous function has completed, will demonstrate
the concurrency race condition.

Sample output:

PS C:\Users\timco\go\src\goroutine> go run .\goroutine.go
71
Done!
PS C:\Users\timco\go\src\goroutine> go run .\goroutine.go
21
Done!
PS C:\Users\timco\go\src\goroutine> go run .\goroutine.go
31
Done!
PS C:\Users\timco\go\src\goroutine> go run .\goroutine.go
61
Done!

**/

package main

import (
	"fmt"
)

func addX() int {
	var x int

	go func() {
		x = 10
	}()

	go func() {
		x = 20
	}()

	go func() {
		x = 30
	}()

	go func() {
		x = 40
	}()

	go func() {
		x = 50
	}()

	go func() {
		x = 60
	}()

	go func() {
		x = 70
	}()

	go func() {
		x = 80
	}()

	return x + 1
}

func main() {

	fmt.Println(addX())

	fmt.Println("Done!")
}
