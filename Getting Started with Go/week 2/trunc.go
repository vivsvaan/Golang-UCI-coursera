package main
import "fmt"

func main() {
	var num float64
	fmt.Print("Enter a float number: ")
	fmt.Scan(&num)
	fmt.Printf("Truncated version is %d", int(num))
}
