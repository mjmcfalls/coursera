package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func sortSlice(s []int, waitgroup *sync.WaitGroup) {

	fmt.Printf("Sorting: %v\n", s)
	sort.Ints(s)
	// fmt.Printf(" ---> Sorted: %v\n", s)
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
	// Test 1 = 85 78 4 10 94 38 55 4 64 100
	// Test 2 = -64 66 20 32 80 -66 -35 -39 -88 -38 -29 43 75 77 -16 34 -99 -50 34
	// test 3 = -88 -92 -8 -34 71 90 -87 71 -68 -49
	var numOfGroups int = 4
	var chunkSize int
	var wg sync.WaitGroup
	var array []int

	// array := []int{85, 78, 4, 10, 94, 38, 55, 4, 64, 100}
	// Get numbers from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter numbers separated by a space: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	// Split numbers into array by space
	stritems := strings.Fields(text)
	// Convert strings to ints, and append to array
	for _, item := range stritems {
		tempInt, _ := strconv.Atoi(item)
		array = append(array, tempInt)
	}

	// Build slice chunksize
	chunkSize = len(array) / numOfGroups

	for i := 0; i < len(array); i += chunkSize {
		// Increment waitgroup
		wg.Add(1)
		// Start goroutine
		go sortSlice(array[i:i+chunkSize], &wg)
	}

	wg.Wait()
	sort.Ints(array)
	fmt.Printf("Sorted numbers: ")
	// fmt.Println("Done")
	for _, value := range array {
		fmt.Printf("%d ", value)
	}

}
