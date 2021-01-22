package main

import "fmt"

func Swap(slice []int, index int) {
	slice[index], slice[index+1] = slice[index+1], slice[index]
}

func BubbleSort(slice []int) {
	for j := 1; j < len(slice); j++ {
		for i := 0; i < len(slice)-j; i++ {
			if slice[i] > slice[i+1] {
				Swap(slice, i)
			}
		}
	}
}

func main() {
	slice := make([]int, 0, 10)
	var input int
	read, _ := fmt.Scan(&input)
	for read > 0 {
		slice = append(slice, input)
		read, _ = fmt.Scan(&input)
	}
	BubbleSort(slice)
	fmt.Print(slice)
}
