package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// read a string from input line
func ReadInputLine() (string, error) {
	stdReader := bufio.NewReader(os.Stdin)
	if inputLine, err := stdReader.ReadString('\n'); err != nil {
		fmt.Println("Invalid Input")
		return "", err
	} else {
		if strings.HasSuffix(inputLine, "\n") {
			inputLine = strings.TrimSuffix(inputLine, "\n")
		}
		return inputLine, nil
	}
}

// convert an array of strings to integers
func StringToIntegers(str string) ([]int, error) {
	str_array := strings.Split(str, " ")
	slice := make([] int, 0, len(str_array))

	for _, v := range str_array {
		if n, err := strconv.Atoi(v); err != nil {
			fmt.Println("Error:", err)
			return nil, err
		} else {

			slice = append(slice, n)
		}
	}

	return slice, nil
}

// read a string input line, return an array intefwea
func ReadIntLine() ([]int, error) {
	if inputString, err := ReadInputLine(); err!=nil {
		return nil, err
	} else {
		return StringToIntegers(inputString)
	}

}

func SwapAdjacent(sequence []int, index int) {
	if index >= len(sequence)-1 {
		return
	}
	temp := sequence[index]
	sequence[index] = sequence[index+1]
	sequence[index+1] = temp
}

func BubbleSort(sequence []int, wg *sync.WaitGroup) {
	fmt.Println("Sequence to be sorted:", sequence)
	lenght := len(sequence)

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght-1-i; j++ {
			if sequence[j] > sequence[j+1] {
				SwapAdjacent(sequence, j)
			}
		}
	}

	if wg != nil {
		wg.Done()
	}
}

func MergeSorted(slice1 []int, slice2 []int, slice12 []int) {

	i, j, k := 0, 0, 0
	for {
		if slice1[i] < slice2[j] {
			slice12[k] = slice1[i]
			k++
			i++
		} else if slice1[i] == slice2[j] {
			slice12[k] = slice1[i]
			k++
			i++

			slice12[k] = slice2[j]
			k++
			j++
		} else if slice1[i] > slice2[j] {
			slice12[k] = slice2[j]
			k++
			j++
		}

		if i == len(slice1) {
			for j < len(slice2) {
				slice12[k] = slice2[j]
				k++
				j++
			}
		}

		if j == len(slice2) {
			for i < len(slice1) {
				slice12[k] = slice1[i]
				k++
				i++
			}
		}

		if k == len(slice12) {
			break
		}
	}
}

func main() {
	fmt.Println("Please enter a sequence of numbers separated by space, then press Enter to stop input:")

	if intSlice, err:= ReadIntLine(); err != nil {
		fmt.Println("Invalid Inputs: Error:", err)
	} else {
		size := len(intSlice)

		// less than 4 numbers, don't need to partition
		if size < 4 {
			BubbleSort(intSlice, nil)
			fmt.Println("Sorted numbers:", intSlice)
		} else {
			// partition into 4 groups
			subSize := size / 4
			slice1 := intSlice[0:subSize]
			slice2 := intSlice[subSize:2*subSize]
			slice3 := intSlice[2*subSize:3*subSize]
			slice4 := intSlice[3*subSize:size]

			// call 4 goroutines to sort concurrently
			var wg sync.WaitGroup
			wg.Add(4)
			go BubbleSort(slice1, &wg)
			go BubbleSort(slice2, &wg)
			go BubbleSort(slice3, &wg)
			go BubbleSort(slice4, &wg)

			wg.Wait()
			//Merge two of them after all finished
			merged_slice1 := make([]int, len(slice1)+len(slice2))
			merged_slice2 := make([]int, len(slice3)+len(slice4))

			MergeSorted(slice1, slice2, merged_slice1)
			MergeSorted(slice3, slice4, merged_slice2)

			// merged above two sorted slices
			finalSorted := make([]int, size)
			MergeSorted(merged_slice1, merged_slice2, finalSorted)

			fmt.Println("Sorted numbers:", finalSorted)
		}
	}

}