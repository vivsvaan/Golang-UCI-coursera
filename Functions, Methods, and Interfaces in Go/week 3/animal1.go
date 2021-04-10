package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food, locomotion, sound string
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.sound)
}

var zoo = map[string]Animal{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

func Input() (string, string) {
	var name, action string
	fmt.Print("> ")
	fmt.Scan(&name, &action)
	return strings.ToLower(name), strings.ToLower(action)
}

func main() {
	for {
		name, action := Input()
		animal, ok := zoo[name]
		if !ok {
			fmt.Println("No such animal")
			continue
		}
		switch action {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Println("Oh, snap!")
		}
	}
}
