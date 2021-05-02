/**
  This program illustrates concurrency and race conditions in Go.

  There are 2 methods that are executed as 2 different goroutines
  The issue here is that one of them prints the value of the global variable x
  while the other loops about a million times to increment it.

  When this program is executed, it prints basically what looks like a random number
  between 0 and 1000000

  This is because both the routines get executed in a concurrent manner. While the long
  loop is executing, the other routine gets scheduled and prints whatever the current value
  of x was when the first routine got interrupted.

  Run this program a few times to make sure that it indeed does print a random looking number.

*/

package main

import (
	"fmt"
	"time"
)

var x int

func main() {
	// Start 2 goroutines
	go start_increment()
	go print_value()

	// Sleep to handle the case where main thread exits and all other go routines are terminated with it.
	time.Sleep(2000 * time.Millisecond)
}

// Go routine setx() increments x up to 100000
func start_increment() {
	for i := 0; i < 100000; i++ {
		x++
	}
}

// Go routine that prints the current value of x
func print_value() {
	fmt.Println(x)
}
