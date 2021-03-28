package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 3)
	var usr_inp string
	var number int

	for {
		fmt.Print("\n Enter integer (or X to quit): ")
		fmt.Scan(&usr_inp)
		if usr_inp == "X" || usr_inp == "x" {
			break
		} else {
			number, _ = strconv.Atoi(usr_inp)
			sli = append(sli, number)
			sort.Ints(sli)
			fmt.Print("Slice is: ")
			for _, val := range sli {
				fmt.Printf("%d ", val)
			}
		}
	}
}
