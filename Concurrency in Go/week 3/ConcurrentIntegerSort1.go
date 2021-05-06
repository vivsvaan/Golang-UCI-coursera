package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {

	var waitGroup sync.WaitGroup
	var sortedIntegers []int

	fmt.Println("Enter integers sperated by a space:")
	fmt.Print("> ")
	integersToSort := getIntegersFromConsoleInput([]int{}, nil)

	sliceLength := len(integersToSort) / 3
	slice1 := integersToSort[0:sliceLength]
	slice2 := integersToSort[sliceLength : sliceLength*2]
	slice3 := integersToSort[sliceLength*2 : sliceLength*3]
	slice4 := integersToSort[sliceLength*3:]

	waitGroup.Add(4)
	go sortIntegerArray(&slice1, &sortedIntegers, &waitGroup)
	go sortIntegerArray(&slice2, &sortedIntegers, &waitGroup)
	go sortIntegerArray(&slice3, &sortedIntegers, &waitGroup)
	go sortIntegerArray(&slice4, &sortedIntegers, &waitGroup)

	waitGroup.Wait()
	sort.Sort(sort.IntSlice(sortedIntegers))
	fmt.Println("Done! Sorted integers:")
	fmt.Println(sortedIntegers)

}

func sortIntegerArray(sliceToSort *[]int, sortedIntegers *[]int, waitGroup *sync.WaitGroup) {
	sort.Sort(sort.IntSlice(*sliceToSort))
	*sortedIntegers = append(*sortedIntegers, *sliceToSort...)
	waitGroup.Done()
}

func getIntegersFromConsoleInput(x []int, err error) []int {
	if err != nil {
		return x
	}
	var d int
	n, err := fmt.Scanf("%d", &d)
	if n == 1 {
		x = append(x, d)
	}
	return getIntegersFromConsoleInput(x, err)
}
