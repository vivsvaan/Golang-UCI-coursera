package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	var file_name string
	names_slice := make([]Name, 0, 1)

	fmt.Print("Enter the file name: ")
	fmt.Scan(&file_name)
	file, err := os.Open(file_name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		name := strings.Split(scanner.Text(), " ")
		names_slice = append(names_slice, Name{Fname: name[0], Lname: name[1]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, name := range names_slice {
		fmt.Printf("First Name: %s \nLast Name: %s\n\n", name.Fname, name.Lname)
	}
}
