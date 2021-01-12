package main

import (
	"fmt"
	"strings"
)

// BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sortedorder.
func BubbleSort(items []int) {
	for i := 0; i < len(items)-1; i++ {
		var swapped bool = false
		for ii := 0; ii < len(items)-1; ii++ {
			if items[ii+1] < items[ii] {
				Swap(items, ii)
				swapped = true
			}
		}
		if !(swapped) {
			break
		}
	}
}

//  Swap() function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice.
func Swap(items []int, index int) {
	temp := items[index]
	items[index] = items[index+1]
	items[index+1] = temp
}

func main() {
	// prompt the user to type in a sequence of up to 10 integers. The program
	// should print the integers out on one line, in sorted order, from least to greatest
	// items := []int{48, 39, 12, 4, 51, 100, 28, 8, 30, 85}
	fmt.Printf("Enter up to 10 numbers separated by a space: ")
	text, _ := reader.ReadString('\n')
	name := strings.TrimSuffix(text, "\n")

	BubbleSort(items)
	fmt.Println(items)
}
