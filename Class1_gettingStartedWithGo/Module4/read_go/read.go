package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
)

// Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. 
// Each field will be a string of size 20 (characters).
type Person struct {
    fname string
    lname string
}

func main(){

	// Write a program which reads information from a file and represents it in a slice of structs. 
	// Assume that there is a text file which contains a series of names. 
	// Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
	var namesSlice []Person
	reader := bufio.NewReader(os.Stdin)

	// Your program should prompt the user for the name of the text file. 
	fmt.Printf("Enter a filename: ")
	rawfile, _ := reader.ReadString('\n')
	filename := strings.TrimSuffix(rawfile, "\n")

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var names []string

	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	file.Close()

	for _, eachline := range names {
		temp := strings.Fields(eachline)
		// Each struct created will be added to a slice
		namesSlice = append(namesSlice, Person{fname:temp[0], lname:temp[1]})
	}
	
	// iterate through your slice of structs and print the first and last names found in each struct.
	for _, v := range namesSlice{
		fmt.Printf("First Name: %s; Last Name: %s\n", v.fname, v.lname)
	}
}