package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now()
	// fmt.Println(now.Local())

	var now time.Time = time.Now()
	fmt.Println(now)

	var utc time.Time = time.Date(2009, time.August, 17, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)

	formatter := "2006-01-02 15:04:05"
	value := "2023-12-19 10:10:10"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error ", err.Error())
	} else {
		fmt.Println(valueTime)
	}
}
