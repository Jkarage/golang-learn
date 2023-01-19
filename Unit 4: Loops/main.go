package main

import "fmt"

var x int

func main() {
	// A general for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// A traditional while loop
	for x < 10 {
		fmt.Println(x)
		x++
	}

	// A traditional infinite loop
	for {
		if x == 11 {
			break
		}
		x++
	}
}
