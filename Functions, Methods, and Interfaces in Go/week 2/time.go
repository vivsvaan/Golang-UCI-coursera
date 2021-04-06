package main

import (
	"fmt"
	"math"
)

func WriteAndRead(s string, y *float64) {
	fmt.Print("Enter ", s, " : ")
	fmt.Scan(y)
}

func GenDisplaceFn(a, s, v float64) func(float64) float64 {
	res := func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v*t + s
	}
	return res
}

func main() {

	var a float64
	var s0 float64
	var v0 float64
	var t float64

	WriteAndRead("Acceleration", &a)
	WriteAndRead("Initial Displacement", &s0)
	WriteAndRead("Initial Velocity", &v0)
	WriteAndRead("Time", &t)

	getTime := GenDisplaceFn(a, s0, v0)
	fmt.Printf("Time : %.2f s", getTime(t))
}
