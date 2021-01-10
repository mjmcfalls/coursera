package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter a floating point number: ")

	for i := 0; i < 2; i++ {
		number, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Cannot read the string.", err)
		}

		number_float, _ := strconv.ParseFloat(strings.TrimSpace(number), 32)
		fmt.Printf("%d\n", int(number_float))
	}
}
