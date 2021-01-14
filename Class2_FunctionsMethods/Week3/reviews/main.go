package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

func (a animal) Eat() {
	fmt.Println("Eats", a.food)
}

func (a animal) Move() {
	fmt.Println("Moves by", a.locomotion)
}

func (a animal) Speak() {
	fmt.Println("The noise that makes is", a.noise)
}

func main() {

	stdin := bufio.NewScanner(os.Stdin)
	var txt string

	var cow, bird, snake animal
	cow.food = "grass"
	cow.locomotion = "walk"
	cow.noise = "moo"

	bird.food = "worms"
	bird.locomotion = "fly"
	bird.noise = "peep"

	snake.food = "mice"
	snake.locomotion = "slither"
	snake.noise = "hsss"

	var a_ptr *animal
	for {
		fmt.Printf("> ")
		if stdin.Scan() {
			txt = stdin.Text()
		}

		tokens := strings.Fields(txt)
		an := strings.ToLower(tokens[0])
		action := strings.ToLower(tokens[1])

		switch an {
		case "cow":
			a_ptr = &cow
		case "bird":
			a_ptr = &bird
		case "snake":
			a_ptr = &snake
		default:
			fmt.Printf("Unregonized animal\n")
			continue
		}

		switch action {
		case "eat":
			a_ptr.Eat()
		case "move":
			a_ptr.Move()
		case "speak":
			a_ptr.Speak()
		default:
			fmt.Printf("Unrecognized action\n")
			continue
		}

	}

}
