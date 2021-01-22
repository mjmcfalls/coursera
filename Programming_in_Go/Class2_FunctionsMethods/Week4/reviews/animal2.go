package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }
type Bird struct{ name string }
type Snake struct{ name string }

func (c *Cow) Eat() {
	fmt.Println("grass")
}

func (c *Cow) Move() {
	fmt.Println("walk")
}

func (c *Cow) Speak() {
	fmt.Println("moo")
}

func (b *Bird) Eat() {
	fmt.Println("worms")
}

func (b *Bird) Move() {
	fmt.Println("fly")
}

func (b *Bird) Speak() {
	fmt.Println("peep")
}

func (s *Snake) Eat() {
	fmt.Println("mice")
}

func (s *Snake) Move() {
	fmt.Println("slither")
}

func (s *Snake) Speak() {
	fmt.Println("hsss")
}

func AcquireInput(scanner *bufio.Scanner) ([]string, error) {
	fmt.Printf("> ")
	scanner.Scan()
	userInput := scanner.Text()

	userInputFields := strings.Fields(userInput)

	if len(userInputFields) != 3 {
		err := errors.New("Invalid number of strings")
		return nil, err
	}

	return userInputFields, nil
}

func CreateAnimal(name, animalType string) (Animal, error) {
	var animal Animal
	var err error = nil

	switch animalType {
	case "cow":
		animal = &Cow{name: name}
	case "bird":
		animal = &Bird{name: name}
	case "snake":
		animal = &Snake{name: name}
	default:
		animal = nil
		err = errors.New("Invalid animal type")
	}

	return animal, err
}

func PerformAction(animal Animal, action string) error {

	var err error = nil

	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		err = errors.New("Invalid action")
	}

	return err
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	namesMap := make(map[string]Animal)

	for {
		inputFields, err := AcquireInput(scanner)
		if err != nil {
			fmt.Println(err)
			continue
		}

		operation := inputFields[0]
		name := inputFields[1]

		switch operation {
		case "newanimal":
			animalType := inputFields[2]
			animal, err := CreateAnimal(name, animalType)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				continue
			}
			fmt.Println("Created it!")
			namesMap[name] = animal

		case "query":
			action := inputFields[2]

			animal, ok := namesMap[name]
			if !ok {
				fmt.Printf("No animal found with name '%s'\n", name)
				continue
			}
			err := PerformAction(animal, action)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				continue
			}

		default:
			fmt.Printf("Error: Invalid command '%s'\n", operation)
		}

	}

}
