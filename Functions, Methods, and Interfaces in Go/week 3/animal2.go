package main

import "fmt"

type Animal struct {
	food string
	locomotion string
	noise string
}

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.locomotion
}

func (animal *Animal) Speak() string {
	return animal.noise
}

func Prompt () (string, string, error) {
	var animal string
	var action string
	fmt.Printf("> ")
	_, err := fmt.Scanf("%s %s", &animal, &action)
	return animal, action, err
}

func CreateAnimal(AnimalType string) Animal {
	switch AnimalType {
	case "cow":
		return Animal{food: "grass", locomotion: "walk", noise: "moo"}
	case "bird":
		return Animal{food: "worms", locomotion: "fly", noise: "peep"}
	case "snake":
		return Animal{food: "mice", locomotion: "slither", noise: "hsss"}
	default:
		return Animal{food: "undefined animal food", locomotion: "undefined animal spoken sound", noise: "undefined animal spoken sound"}
	}
}

func GetAnimalInfo (AnimalType string, AnimalInfo string) (string) {
	animal := CreateAnimal(AnimalType)
	switch AnimalInfo {
	case "eat":
		return animal.Eat()
	case "move":
		return animal.Move()
	case "speak":
		return animal.Speak()
	default:
		return "undefined animal information data"
	}
}

func main() {
	for {
		animal, action, err := Prompt()
		if err != nil {
			fmt.Printf("Error during readind data: %v", err)
			break
		}

		fmt.Println(GetAnimalInfo(animal, action))
	}
}

