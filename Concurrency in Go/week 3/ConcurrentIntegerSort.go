package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func StringToIntSlice(str string, sep string) []int {
	inputSplit := strings.Split(str, sep)
	intSlc := make([]int, 0, len(inputSplit))

	for _, v := range inputSplit {
		n, err := strconv.Atoi(v)

		if err == nil {
			intSlc = append(intSlc, n)
		}
	}
	return intSlc
}

func MakeNSubarrays(array []int, n float32) [][]int {
	rem := float32(0.0)
	return_array := make([][]int, 0)
	avg := float32(len(array)) / n

	for int(rem) < len(array) {
		return_array = append(return_array, array[int(rem):int(rem+avg)])
		rem = rem + avg
	}
	return return_array

}

func SortArray(slice []int, channel chan []int, wg *sync.WaitGroup) {
	sort.Ints(slice)
	fmt.Println("Sorted slice is: ", slice)
	wg.Done()
	channel <- slice
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan []int, 4)

	fmt.Println("Enter a series of integers.")
	fmt.Print(">")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	array := StringToIntSlice(input, " ")

	fmt.Println("Entered slice is: ", array)

	divided_array := MakeNSubarrays(array, 4)

	wg.Add(4)
	go SortArray(divided_array[0], channel, &wg)
	go SortArray(divided_array[1], channel, &wg)
	go SortArray(divided_array[2], channel, &wg)
	go SortArray(divided_array[3], channel, &wg)
	slice1 := <-channel
	slice2 := <-channel
	slice3 := <-channel
	slice4 := <-channel

	wg.Wait()

	final_array := []int{}
	final_array = append(slice1, slice2...)
	final_array = append(final_array, slice3...)
	final_array = append(final_array, slice4...)

	sort.Ints(final_array)

	fmt.Println("Sorted array is: ", final_array)

}
