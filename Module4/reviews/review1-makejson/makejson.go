package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	var name string
	fmt.Printf("Enter name:")
	fmt.Scanln(&name)
	fmt.Printf("Enter address:")
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	// Remove end of line character
	address := line[:len(line)-1]
	m := map[string]string{
		"name":    name,
		"address": address,
	}
	j, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(j))
}
