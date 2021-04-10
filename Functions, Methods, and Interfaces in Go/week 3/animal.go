package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var animal, action string

	animalMap := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		fmt.Print("> ")
		fmt.Scanln(&animal, &action)

		switch action {
		case "eat":
			f := animalMap[animal]
			f.Eat()
		case "move":
			f := animalMap[animal]
			f.Move()
		case "speak":
			f := animalMap[animal]
			f.Speak()
		default:
			fmt.Println("Oh, snap!")
		}

	}

}
