package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type bird struct {
	food       string
	locomotion string
	noise      string
}

type cow struct {
	food       string
	locomotion string
	noise      string
}

type snake struct {
	food       string
	locomotion string
	noise      string
}
type Animal interface {
	Eat()
	Move()
	Speak()
}

func (a *bird) Eat() {
	fmt.Printf("Eats: %s\n", a.food)
}

func (a *bird) Move() {
	fmt.Printf("Move: %s\n", a.locomotion)
}

func (a *bird) Speak() {
	fmt.Printf("Speak: %s\n", a.noise)
}

func (a *cow) Eat() {
	fmt.Printf("Eats: %s\n", a.food)
}

func (a *cow) Move() {
	fmt.Printf("Move: %s\n", a.locomotion)
}

func (a *cow) Speak() {
	fmt.Printf("Speak: %s\n", a.noise)
}

func (a *snake) Eat() {
	fmt.Printf("Eats: %s\n", a.food)
}

func (a *snake) Move() {
	fmt.Printf("Move: %s\n", a.locomotion)
}

func (a *snake) Speak() {
	fmt.Printf("Speak: %s\n", a.noise)
}

func Eat(a *Animal) {
	iface := *a
	iface.Eat()
}

func Move(a *Animal) {
	iface := *a
	iface.Move()
}
func Speak(a *Animal) {
	iface := *a
	iface.Speak()
}

func main() {
	animals := make(map[string]Animal)
	var a Animal

	animals["cow"] = &cow{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = &bird{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = &snake{food: "mice", locomotion: "slither", noise: "hsss"}

	// Each “newanimal” command must be a single line containing three strings.
	// The first string is “newanimal”. The second string is an arbitrary string which will be the name of the new animal.
	// The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
	// Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

	// Each “query” command must be a single line containing 3 strings. The first string is “query”.
	// The second string is the name of the animal.
	// The third string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
	// Your program should process each query command by printing out the requested data.

	// Define an interface type called Animal which describes the methods of an animal.
	// Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values.
	// The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion,
	// and the Speak() method should print the animal’s spoken sound.
	// Define three types Cow, Bird, and Snake.
	// For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
	// When the user creates an animal, create an object of the appropriate type.
	// Your program should call the appropriate method when the user issues a query command.

	fmt.Printf("%s\n", "Query: query animal eat,move,speak")
	fmt.Printf("%s\n", "Add: newanimal animal type")
	fmt.Printf("%s\n", "Exit: x")

	reader := bufio.NewReader(os.Stdin)
Loop:
	for {

		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		stritems := strings.Fields(strings.ToLower(text))
		if len(stritems) == 3 {
			// var tempanimal *animal
			// a = tempanimal
			if stritems[0] == "query" {
				a = animals[stritems[1]]
				// GetAnimalAction(&a)
				switch stritems[2] {
				case "eat":
					Eat(&a)
				case "move":
					// a = animals[stritems[1]]
					Move(&a)
				case "speak":
					// a = animals[stritems[1]]
					Speak(&a)
				case "x":
					fmt.Println("Exiting")
					break Loop
				default:
					fmt.Println("Unknown action")
				}
			} else if stritems[0] == "newanimal" {
				switch stritems[2] {
				case "bird":
					animals[stritems[1]] = &bird{food: "worms", locomotion: "fly", noise: "peep"}
					fmt.Println("Created it!")
				case "snake":
					animals[stritems[1]] = &snake{food: "mice", locomotion: "slither", noise: "hsss"}
					fmt.Println("Created it!")
				case "cow":
					animals[stritems[1]] = &cow{food: "grass", locomotion: "walk", noise: "moo"}
					fmt.Println("Created it!")
				}

			}
		} else if len(stritems) == 1 {
			switch stritems[0] {
			case "x":
				fmt.Println("Exiting")
				break Loop
			default:
				fmt.Println("Invalid Entry!!")
				fmt.Printf("%s\n", "Query: query animal eat,move,speak")
				fmt.Printf("%s\n", "Add: newanimal animal type")
				fmt.Printf("%s\n", "Exit: x")
			}
		} else {
			fmt.Println("Invalid Entry!!")
			fmt.Printf("%s\n", "Query: query animal eat,move,speak")
			fmt.Printf("%s\n", "Add: newanimal animal type")
			fmt.Printf("%s\n", "Exit: x")
		}
	}
}
