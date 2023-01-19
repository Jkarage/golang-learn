package main

import "fmt"

func main() {
	// Declaring numbers using var keyword
	var age int = 40
	var bankBallance float64 = 5678.874

	// Using short-declaration syntax{ Alloed inside functions only }
	fullName := "Joseph Karage"
	hasCredit := true

	// Declaring an array { My friends are immutable }
	friends := [5]string{"RoseMary", "Khadja", "Fadhili", "Charles", "Caroline"}

	// Declaring a slice { Enemies are mutable}
	enemies := []string{"Wallace", "Mendez"}

	// Declaring a map
	ages := make(map[string]int)
	fmt.Println(age, bankBallance, fullName, hasCredit, friends, enemies, ages)

	// Declaring a struct
	type Phone struct {
		phone       int
		countryCode int
	}

	type Profile struct {
		username string
		age      int
		email    string
		Phone
	}
}
