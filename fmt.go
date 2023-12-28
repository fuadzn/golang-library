package main

import "fmt"

func main() {
	firstName := "Fuad"
	lastName := "Zaidan"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
