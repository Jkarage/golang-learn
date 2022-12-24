package main

import "fmt"

func main() {
	// Declaring numbers using var keyword
	var age int = 40
	var bankBallance float64 = 5678.874

	// Using short-declaration syntax{ Alloed inside functions only }
	fullName := "Joseph Karage"
	isCreditable := true

	// Declaring an array
	friends := [5]string{"RoseMary", "Khadja", "Fadhili", "Charles", "Caroline"}

	// Declaring a slice
	enemies := []string{"Wallace", "Mendez"}

	fmt.Println(age, bankBallance, fullName, isCreditable, friends, enemies)
}
