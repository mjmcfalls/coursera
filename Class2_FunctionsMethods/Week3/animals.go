package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func Eat(a Animal) {
	fmt.Printf("Eats: %s\n", a.food)
}

func Move(a Animal) {
	fmt.Printf("Move: %s\n", a.locomotion)
}

func Speak(a Animal) {
	fmt.Printf("Speak: %s\n", a.noise)
}

func main() {
	animals := make(map[string]Animal)
	animals["cow"] = Animal{food: "grass", locomotion: "walk", noise: "moo"}
	animals["bird"] = Animal{food: "worms", locomotion: "fly", noise: "peep"}
	animals["snake"] = Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	// program should present the user with a prompt, “>”, to indicate that the user can type a request.
	// Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
	// Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
	// The first string is the name of an animal, either “cow”, “bird”, or “snake”.
	// The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
	// Your program should process each request by printing out the requested data.

	// program should present the user with a prompt, “>”, to indicate that the user can type a request.
	// Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
	// Your program should continue in this loop forever. Every request from the user must be a single line containing 2 strings.
	// The first string is the name of an animal, either “cow”, “bird”, or “snake”.
	// The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”.
	// Your program should process each request by printing out the requested data.
	reader := bufio.NewReader(os.Stdin)
Loop:
	for {
		fmt.Printf("> ")
		text, _ := reader.ReadString('\n')
		stritems := strings.Fields(strings.ToLower(text))
		if len(stritems) > 1 {
			switch stritems[1] {
			case "eat":
				Eat(animals[stritems[0]])
			case "move":
				Move(animals[stritems[0]])
			case "speak":
				Speak(animals[stritems[0]])
			case "x":
				fmt.Println("Exiting")
				break Loop
			default:
				fmt.Println("Unknown action")
			}
		} else {
			switch stritems[0] {
			case "x":
				fmt.Println("Exiting")
				break Loop
			default:
				fmt.Println("Unknown action")
			}
		}

	}

}
