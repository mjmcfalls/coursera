package main

import "fmt"

// BubbleSort() which takes a slice of integers as an argument and returns nothing.
// The BubbleSort() function should modify the slice so that the elements are in sortedorder.
func BubbleSort(items []int) {

	for i := 0; i < len(items)-1; i++ {
		// fmt.Printf("i: %d, i+1: %d\n", i, i+1)
		var swapped bool = false
		for ii := 0; ii < len(items)-1; ii++ {
			// fmt.Printf("i: %d,v: %d, i+1: %d, v+1: %d\n", i, items[i], i+1, items[i+1])
			// fmt.Print(items[i : i+2])
			if items[ii+1] < items[ii] {
				// fmt.Printf("PreSwap: %d,%d\n", items[ii], items[ii+1])
				Swap(items, ii)
				// fmt.Printf("PostSwap: %d,%d\n", items[i], items[i+1])
				swapped = true
			}

		}
		// fmt.Println(items)
		if !(swapped) {
			break
		}

	}

	// fmt.Printf("%d,%d\n", subitems[0], subitems[1])

}

//  Swap() function should take two arguments, a slice of integers and an index value i which
// indicates a position in the slice.
func Swap(items []int, index int) {
	// fmt.Printf("Preswaps: %d,%d\n", items[index], items[index+1])
	temp := items[index]
	items[index] = items[index+1]
	items[index+1] = temp
	// fmt.Printf("%d,%d\n", x, subitems[1])
}

func main() {
	// prompt the user to type in a sequence of up to 10 integers. The program
	// should print the integers out on one line, in sorted order, from least to greatest
	items := []int{48, 39, 12, 4, 51, 100, 28, 8, 30, 85}

	BubbleSort(items)
	fmt.Println(items)
}
