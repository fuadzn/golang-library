package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"fuad", "Zaidan", "najib"}
	values := []int{2, 4, 5, 6, 7}

	fmt.Println(slices.Min(names))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Max(values))
}
