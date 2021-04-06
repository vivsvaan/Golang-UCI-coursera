package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	/** Formula for displacement s as a function of time t,
	* acceleration a, initial velocity vo, and initial displacement so.
	* s = ½ a t2 + vot + so
	 */

	disp := func(t float64) float64 {
		return (((a * t * t) / 2) + (vo * t) + so)
	}

	return disp

}

func main() {
	var acc, ini_vel, ini_disp, time float64
	var err1, err2, err3 error

	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter values for accelaration, initial velocity and initial displacement (space-separated).... \n")

	input, _ := consoleReader.ReadString('\n')
	values := strings.Split(input, " ")

	acc, err1 = strconv.ParseFloat(values[0], 64)
	if err1 != nil {
		fmt.Println("Invalid accelaration! Exiting Program...")
		return
	}

	ini_vel, err2 = strconv.ParseFloat(values[1], 64)
	if err2 != nil {
		fmt.Println("Invalid velocity! Exiting Program...")
		return
	}

	ini_disp, err3 = strconv.ParseFloat(values[2], 64)
	if err3 != nil {
		fmt.Println("Invalid displacement! Exiting Program...")
		print(err3)
		return
	}

	fn := GenDisplaceFn(acc, ini_vel, ini_disp)

	fmt.Print("\nEnter time: ")
	fmt.Scan(&time)

	fmt.Print(fn(time))

}
