package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	ListSize int    = 10
	QuitCode string = "X"
)

func Swap(numbers []int, i int) {
	tmp := numbers[i]
	numbers[i] = numbers[i+1]
	numbers[i+1] = tmp
}

func BubbleSort(numbers []int) {
	Swapped := true
	for Swapped {
		Swapped = false
		for i := 0; i < len(numbers)-1; i++ {
			if numbers[i] > numbers[i+1] {
				Swap(numbers, i)
				Swapped = true
			}
		}
	}
}

func stripSpaces(input string) string {
	re_leadclose_whtsp := regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	re_inside_whtsp := regexp.MustCompile(`[\s\p{Zs}]{2,}`)
	result := re_leadclose_whtsp.ReplaceAllString(input, "")
	result = re_inside_whtsp.ReplaceAllString(result, " ")
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
	fmt.Println("Enter X to exit...")
	fmt.Printf("Please enter a list of integers (space-separated): ")

	for scanner.Scan() {
		var strVal string = scanner.Text()
		if strVal == QuitCode {
			os.Exit(0)
		}
		strVal = stripSpaces(strVal)
		var numlist = strings.Split(strVal, " ")
		var intSlice = make([]int, 0)
		for _, i := range numlist {
			j, err := strconv.Atoi(i)
			if err != nil {
				fmt.Println("Invalid input! Please try again...")
				fmt.Printf("Please enter a list of integers (space-separated): ")
				continue
			}
			intSlice = append(intSlice, j)
		}
		fmt.Printf("len(intSlice): %d; cap(intSlice): %d\n", len(intSlice), cap(intSlice))
		if len(intSlice) > ListSize {
			fmt.Printf("Please enter a maximum of %d integers\n", ListSize)
			fmt.Printf("Please enter a list of integers (space-separated): ")
			continue
		}
		BubbleSort(intSlice)
		fmt.Println(intSlice)
		fmt.Printf("Please enter a list of integers (space-separated): ")
	}
}
