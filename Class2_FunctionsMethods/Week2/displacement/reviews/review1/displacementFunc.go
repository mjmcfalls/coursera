package main

import "fmt"

func GenDisplaceFn(a, v, s float64) func(float642 float64) float64 {
	return func(time float64) float64 {
		return a*time*time/2 + v*time + s
	}
}

func main() {
	var a, v, s float64
	_, _ = fmt.Scan(&a, &v, &s)
	fn := GenDisplaceFn(a, v, s)
	var time float64
	_, _ = fmt.Scan(&time)
	fmt.Print(fn(time))
}
