package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name    string
	Address string
}

func main() {

	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name .... \n")

	name, _ := consoleReader.ReadString('\n')
	name = name[:len(name)-1]

	fmt.Print("Enter your adress .... \n")
	addr, _ := consoleReader.ReadString('\n')
	addr = addr[:len(addr)-1]

	p1 := Person{Name: name, Address: addr}

	jsonData, _ := json.Marshal(p1)
	fmt.Println(string(jsonData))
}
