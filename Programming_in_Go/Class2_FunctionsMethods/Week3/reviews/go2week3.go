package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	eat, move, speak string
}

func (input Animal) Eat() {
	fmt.Print(input.eat)
}

func (input Animal) Move() {
	fmt.Print(input.move)
}

func (input Animal) Speak() {
	fmt.Print(input.speak)
}

var cow = Animal{"grass", "walk", "moo"}
var bird = Animal{"worms", "fly", "peep"}
var snake = Animal{"mice", "slither", "hsss"}

var aniMap = map[string]Animal {
	"cow": cow,
	"bird": bird,
	"snake": snake,
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	for {
		fmt.Println()
		fmt.Println("Enter an animal [cow,bird,snake] followed by an action [eat,move,speak] on the same line.")
		fmt.Print("> ")
		input, _ := scanner.ReadString('\n')
		parsedInput := strings.Fields(input)
		fmt.Println(input)
		inAnimal := parsedInput[0]
		inAction := parsedInput[1]
		fmt.Println("Animal: ", parsedInput[0])
		fmt.Println("Action: ", parsedInput[1])
		aniData := aniMap[inAnimal]
		fmt.Print("Result: ")
		if inAction == "eat"{
			// call eat
			aniData.Eat()
		} else if inAction == "move" {
			// call move
			aniData.Move()
		} else if inAction == "speak" {
			// call speak
			aniData.Speak()
		}
		fmt.Println()
	}
}