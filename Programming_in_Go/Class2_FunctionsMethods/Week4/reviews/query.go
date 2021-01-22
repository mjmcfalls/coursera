package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

// define interface
type Animal interface {
	Eat() 
	Move()
	Speak()
}

// define types
type Cow struct {
	food, locomotion, sound string
}

type Bird struct {
	food, locomotion, sound string
}

type Snake struct {
	food, locomotion, sound string
}

// satisfy Animal interface for Cow type
func (c Cow) Eat() {
    fmt.Println(c.food)
}
func (c Cow) Move() {
    fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
    fmt.Println(c.sound)
}

// satisfy Animal interface for Bird type
func (b Bird) Eat() {
    fmt.Println(b.food)
}
func (b Bird) Move() {
    fmt.Println(b.locomotion)
}
func (b Bird) Speak() {
    fmt.Println(b.sound)
}

// satisfy Animal interface for Snake type
func (s Snake) Eat() {
    fmt.Println(s.food)
}
func (s Snake) Move() {
    fmt.Println(s.locomotion)
}
func (s Snake) Speak() {
    fmt.Println(s.sound)
}

func main() {

	m := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		line := strings.Split(scanner.Text(), " ")

		command := line[0]
		name := line[1] // name
		param := line[2] // type or info

		if command == "newanimal" {
			if param == "cow" {
				m[name] = Cow{"grass", "walk", "moo"}
			}else if param == "bird" {
				m[name] = Bird{"worms", "fly", "peep"}
			}else if param == "snake" {
				m[name] = Snake{"mice", "slither", "hsss"}
			}

			fmt.Println("Created it!")
		}else if command == "query" {
			if param == "eat" {
				m[name].Eat()
			} else if param == "move" {
				m[name].Move()
			} else if param == "speak" {
				m[name].Speak()
			}
		}
	}
}

/*
Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”. 
The second string is an arbitrary string which will be the name of the new animal. 
The third string is the type of the new animal, either “cow”, “bird”, or “snake”. 
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”. 
The second string is the name of the animal. The third string is the name of the information requested about the animal,
either “eat”, “move”, or “speak”. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. 
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound. 
Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. 
When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.
*/	

