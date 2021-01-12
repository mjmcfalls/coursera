package main
import (
	"fmt"
	"sort"
	"strconv"
)

func main(){
	intslice := make([]int, 0, 3)
	var input string

	for input != "X" {
		fmt.Printf("Enter a number or X to exit:")
		fmt.Scan(&input)
		// Append to slice
		if input == "X"{
			continue
		} else{
			in, _ := strconv.Atoi(input)
			intslice = append(intslice, in)
			// sort slice
			sort.Ints(intslice)
			// show slice
			fmt.Println(intslice)
		}

	}
}