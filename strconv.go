package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(boolean)
	}

	result, err := strconv.Atoi("1000")
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(result)
	}
}
