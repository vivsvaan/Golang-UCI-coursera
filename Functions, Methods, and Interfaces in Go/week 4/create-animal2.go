/*
The goal of this assignment is to write a GoLang routine that allows
users to create a set of animals and then get information about those animals.

Your program should present the user with a prompt, “>”, to indicate that the user can type a request.
Your program should accept one command at a time from the user, print out a response,
and print out a new prompt on a new line. Your program should continue in this loop forever.
Every command from the user must be either a “newanimal” command or a “query” command.

Each “newanimal” command must be a single line containing three strings. The first string is “newanimal”.
The second string is an arbitrary string which will be the name of the new animal.
The third string is the type of the new animal, either “cow”, “bird”, or “snake”.
Your program should process each newanimal command by creating the new animal and printing “Created it!” on the screen.

Each “query” command must be a single line containing 3 strings. The first string is “query”.
The second string is the name of the animal. The third string is the name of the information requested
about the animal, either “eat”, “move”, or “speak”. Your program should process each query command
by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal.
Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(),
which take no arguments and return no values. The Eat() method should print the animal’s food,
the Move() method should print the animal’s locomotion, and the Speak() method should print the animal’s spoken sound.
Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(),
and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface.
When the user creates an animal, create an object of the appropriate type.
Your program should call the appropriate method when the user issues a query command.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	argCount      int    = 3
	commandPrompt string = "> "
	quitCode      string = "X"
)

var animalMap map[string]Animal
var animalArray = []string{"cow", "bird", "snake"}
var requestArray = []string{"eat", "move", "speak"}
var commandArray = []string{"newanimal", "query", "show"}

func specifyUsage() {
	fmt.Println("Usage 1: newanimal <animal_name> <animal_type>")
	fmt.Println("         where animal_name is a string")
	fmt.Printf("         where animal_type is one of (%s)\n\n", strings.Join(animalArray, ","))
	fmt.Println("Usage 2: query <animal_name> <request>")
	fmt.Println("         where animal_name is a string")
	fmt.Printf("         where request is one of (%s)\n\n", strings.Join(requestArray, ","))
	fmt.Println("Usage 3: show ")
	fmt.Println("         Lists all the animals entered")
	fmt.Println("Enter X to exit...")
}

func init() {
	// https://golang.org/doc/effective_go.html#allocation_new
	animalMap = make(map[string]Animal)
}

func show() {
	for k, v := range animalMap {
		fmt.Printf("%s is a ", k)
		switch v.(type) {
		case Cow:
			fmt.Print("Cow ")
		case Bird:
			fmt.Print("Bird ")
		case Snake:
			fmt.Print("Snake ")
		}
		fmt.Println(v)
	}
}

func createNewAnimal(animalName, animalType string) {
	switch animalType {
	case animalArray[0]:
		animal := Cow{food: "grass", locomotion: "walk", noise: "moo"}
		animalMap[animalName] = animal
	case animalArray[1]:
		animal := Bird{food: "worms", locomotion: "fly", noise: "peep"}
		animalMap[animalName] = animal
	case animalArray[2]:
		animal := Snake{food: "mice", locomotion: "slither", noise: "hsss"}
		animalMap[animalName] = animal
	default:
		fmt.Errorf("%s: Invalid animal type", animalType)
		specifyUsage()
		return
	}
	fmt.Println("Created it!")
}

func queryAnimal(animalName, info string) {
	animal := animalMap[animalName]
	switch info {
	case requestArray[0]:
		fmt.Printf("%s eats %s\n", animalName, animal.Eat())
	case requestArray[1]:
		fmt.Printf("%s moves by %s\n", animalName, animal.Move())
	case requestArray[2]:
		fmt.Printf("%s speaks %s\n", animalName, animal.Speak())
	default:
		fmt.Errorf("%s: Invalid Request", info)
		specifyUsage()
	}
}

func parseCommand(list []string) {
	if len(list) != argCount {
		fmt.Errorf("Usage error: require %d arguments", argCount)
	}
	switch list[0] {
	case commandArray[0]:
		if itemIsInArray(list[2], animalArray) {
			createNewAnimal(list[1], list[2])
		}
	case commandArray[1]:
		if itemIsInArray(list[2], requestArray) {
			queryAnimal(list[1], list[2])
		}
	case commandArray[2]:
		show()
	}
	fmt.Print(commandPrompt)
}

func main()  {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Print(err)
	}
	specifyUsage()
	fmt.Println(commandPrompt)
	for scanner.Scan() {
		userInput := scanner.Text()
		if userInput == quitCode {
			os.Exit(0)
		}
		userInput = stripSpaces(userInput)
		userInput = strings.ToLower(userInput)
		var list = strings.Split(userInput, " ")

		if !(len(list) == argCount && itemIsInArray(list[0], commandArray)) {
			if list[0] != commandArray[2] {
				specifyUsage()
				fmt.Print(commandPrompt)
				continue
			}
		}
		parseCommand(list)
	}
}

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct{ food, locomotion, noise string}
type Bird struct{ food, locomotion, noise string}
type Snake struct{ food, locomotion, noise string}

func (a Cow) Eat() string {
	return a.food
}
func (a Cow) Move() string {
	return a.locomotion
}
func (a Cow) Speak() string {
	return a.noise
}

func (a Bird) Eat() string {
	return a.food
}

func (a Bird) Move() string {
	return a.locomotion
}

func (a Bird) Speak() string {
	return a.noise
}

func (a Snake) Eat() string {
	return a.food
}

func (a Snake) Move() string {
	return a.locomotion
}

func (a Snake) Speak() string {
	return a.noise
}


func stripSpaces(input string) string {
	reLeadcloseWhtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	reInsideWhtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := reLeadcloseWhtsp.ReplaceAllString(input, "")
	result = reInsideWhtsp.ReplaceAllString(result, " ")
	return result
}

func itemIsInArray(item string, array []string) bool {
	for _, i := range array {
		if i == item {
			return true
		}
	}
	return false
}