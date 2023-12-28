package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Fuad Zaidan", "Fuad"))
	fmt.Println(strings.Split("Fuad Zaidan", " "))
	fmt.Println(strings.ToLower("Fuad Zaidan"))
	fmt.Println(strings.ToUpper("Fuad Zaidan"))
	fmt.Println(strings.Trim("   Fuad Zaidan", " "))
}
