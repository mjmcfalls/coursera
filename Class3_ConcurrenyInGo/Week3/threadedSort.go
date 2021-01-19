package main

import (
	"fmt"
	"sort"
	"sync"
)

func sortSlice(s []int, waitgroup *sync.WaitGroup) {

	fmt.Printf("PreSort: %v\n", s)
	sort.Ints(s)
	fmt.Printf("Sorted: %v\n", s)
	waitgroup.Done()

}

func main() {
	// Write a program to sort an array of integers.
	// The program should partition the array into 4 parts, each of which is sorted by a different goroutine.
	// Each partition should be of approximately equal size.
	// Then the main goroutine should merge the 4 sorted subarrays into one large sorted array.

	// The program should prompt the user to input a series of integers.
	// Each goroutine which sorts ¼ of the array should print the subarray that it will sort.
	// When sorting is complete, the main goroutine should print the entire sorted list.
	// 85	78	4	10	94	38	55	4	64	100
	// 28	93	41	23	81	70	27	63	82	59
	var numOfGroups int = 4
	var chunkSize int
	var count int = 1
	var wg sync.WaitGroup

	array := [10]int{85, 78, 4, 10, 94, 38, 55, 4, 64, 100}
	chunkSize = len(array) / numOfGroups

	fmt.Printf("%v\n", array)
	fmt.Printf("Length: %v; items per quarter: %v\n", len(array), len(array)/chunkSize)
	for i := 0; i < len(array); i += chunkSize {
		if i+chunkSize > len(array) {
			wg.Add(count)
			// fmt.Println("i > chunk: %v", i)
			sortSlice(array[i:len(array)], &wg)
			// fmt.Printf("Sorting: %v\n", array[i:len(array)])

		} else {
			wg.Add(count)
			// fmt.Println("i: %v", i)
			sortSlice(array[i:i+chunkSize], &wg)
			// fmt.Printf("Sorting: %v\n", array[i:i+chunkSize])
		}
		count++

	}

	fmt.Println("Done")

}
