/*
Question:
	Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line. 

	Your program will define a name struct which has two fields, fname for the first name, and lname for the last name. Each field will be a string of size 20 (characters).

	Your program should prompt the user for the name of the text file. Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file. Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/
package main

import (
	"strings"
    "bufio"
    "fmt"
    "log"
    "os"
)

type Name struct{
	fname string
	lname string
}

func main() {
	
	var filename string
	fmt.Printf("Please Input FileName: ")
	fmt.Scan(&filename)

    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	
	result := make([]Name, 0, 10)
    for scanner.Scan() {
		divides := strings.Split(scanner.Text(), " ")
		n1 := Name{fname: divides[0], lname: divides[1]}
		result = append(result, n1)
    }
	
	fmt.Println()
	fmt.Println("Result:")
	
	for _, name := range result {
		fmt.Println("  First Name: ", name.fname)
		fmt.Println("  Last Name: ", name.lname)
		fmt.Println()
	}
}