package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, index int) {
	temp := slice[index]
	slice[index] = slice[index+1]
	slice[index+1] = temp
}

func BubbleSort(slice []int) {
	n := len(slice)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

func main() {
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a list of space separated integers: ")

	input, _ := consoleReader.ReadString('\n')
	input = input[:len(input)-1]

	var numbers = strings.Split(input, " ")

	var slice = make([]int, 0)

	for _, i := range numbers {
		j, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Invalid input! Exiting Program...")
			return
		}
		slice = append(slice, j)
	}
	if len(slice) > 10 {
		fmt.Println("Please enter a maximum of 10 integers. Exiting Program...")
		return
	}

	BubbleSort(slice)

	fmt.Println("Sorted Array is ")

	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
}
