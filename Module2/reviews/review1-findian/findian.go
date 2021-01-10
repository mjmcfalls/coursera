package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Get value from user and compute the lower case string
	// Scan() is not suitable to get a string with spaces
	fmt.Printf("Enter a string:\t")
	reader := bufio.NewReader(os.Stdin)
	bytes, _, _ := reader.ReadLine()
	value := string(bytes)
	value_lower := strings.ToLower(value)

	// Determine if lower case string is "findian"
	if strings.HasPrefix(value_lower, "i") &&
		strings.HasSuffix(value_lower, "n") &&
		strings.Contains(value_lower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
