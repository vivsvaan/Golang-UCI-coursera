package main

import (
	"fmt"
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

	fmt.Print("Enter Accelaration: ")
	fmt.Scan(&acc)
	fmt.Print("Enter initial Velocity: ")
	fmt.Scan(&ini_vel)
	fmt.Print("Enter initial Displacement: ")
	fmt.Scan(&ini_disp)

	fn := GenDisplaceFn(acc, ini_vel, ini_disp)

	fmt.Print("\nEnter time: ")
	fmt.Scan(&time)

	fmt.Printf("Time : %.2f s", fn(time))

}
