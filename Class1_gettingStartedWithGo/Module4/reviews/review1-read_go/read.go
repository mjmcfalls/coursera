package main

import (
	"fmt"
	"os"
	"strings"
)

//Person struct with first and last name
type Person struct {
	fname, lname string
}

func main() {
	var filename string
	fmt.Printf("Enter filename: ")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("error:", err)
	}
	var persons []Person

	prevData := ""
	fName, lName := "", ""
	look4FirstName := true
	for true {
		data := make([]byte, 20)
		count, _ := file.Read(data)
		if count == 0 {
			break
		}

		str := prevData + string(data)
		curPos := 0
		for true {
			if look4FirstName {
				spaceIdx := strings.Index(str[curPos:], " ")
				if spaceIdx == -1 {
					break
				}
				fName = str[curPos : curPos+spaceIdx]
				look4FirstName = false
				curPos = curPos + spaceIdx + 1
			} else {
				newLineIdx := strings.Index(str[curPos:], "\n")
				if newLineIdx == -1 {
					break
				}
				lName = str[curPos : curPos+newLineIdx]
				look4FirstName = true
				curPos = curPos + newLineIdx + 1
				p := Person{fName, lName}
				persons = append(persons, p)
			}
		}
		prevData = str[curPos:]
	}
	for _, p := range persons {
		fmt.Println(p)
	}
}
