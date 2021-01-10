package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	var foundstr string = "Found!"
	var notfound string = "Not Found!"

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a string: ")
	text, _ := reader.ReadString('\n')
	lowertxt := strings.ToLower(strings.TrimSuffix(text, "\n"))
	// Starts with i, ends with n contains a == Found!
	// else == Not Found!
	// Found! tests: "ian", "Ian", "iuiygaygn", "I d skd a efju N" 
	// Not Found tests: ihhhhhn, ina, xian. 
	if strings.Contains(lowertxt,"a") && strings.Contains(lowertxt,"i") && strings.Contains(lowertxt,"n") {
		if strings.HasPrefix(lowertxt, "i") && strings.HasSuffix(lowertxt, "n") {
			fmt.Printf("%s\n", foundstr)
		} else {
			fmt.Printf("%s\n", notfound)
		}

	} else {
		fmt.Printf("%s\n", notfound)
	}
}