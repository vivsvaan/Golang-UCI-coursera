package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food, locomotion, noise string
}

type Bird struct {
	food, locomotion, noise string
}

type Snake struct {
	food, locomotion, noise string
}

func (animal *Cow) Eat() {
	fmt.Println(animal.food)
}

func (animal *Cow) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Cow) Speak() {
	fmt.Println(animal.noise)
}

func (animal *Bird) Eat() {
	fmt.Println(animal.food)
}

func (animal *Bird) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Bird) Speak() {
	fmt.Println(animal.noise)
}

func (animal *Snake) Eat() {
	fmt.Println(animal.food)
}

func (animal *Snake) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Snake) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var req_type, animal_name, action string

	animalTypeMap := map[string]string{
		"cow":   "cow",
		"bird":  "bird",
		"snake": "snake",
	}

	cowMap := map[string]Cow{
		"cow": {"grass", "walk", "moo"},
	}

	birdMap := map[string]Bird{
		"bird": {"worms", "fly", "peep"},
	}

	snakeMap := map[string]Snake{
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		fmt.Print("> ")
		fmt.Scanln(&req_type, &animal_name, &action)

		switch req_type {
		case "newanimal":
			switch action {
			case "cow":
				cowMap[animal_name] = Cow{"grass", "walk", "moo"}
				animalTypeMap[animal_name] = "cow"
				fmt.Println("Created It!")
			case "bird":
				birdMap[animal_name] = Bird{"worms", "fly", "peep"}
				animalTypeMap[animal_name] = "bird"
				fmt.Println("Created It!")
			case "snake":
				snakeMap[animal_name] = Snake{"mice", "slither", "hsss"}
				animalTypeMap[animal_name] = "snake"
				fmt.Println("Created It!")
			}
		case "query":
			animal_type := animalTypeMap[animal_name]

			switch animal_type {
			case "cow":
				animal := cowMap[animal_name]
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
			case "bird":
				animal := birdMap[animal_name]
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
			case "snake":
				animal := snakeMap[animal_name]
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
	}
}
