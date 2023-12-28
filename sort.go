package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp
}

func main() {
	users := []User{
		{"Eko", 25},
		{"Budi", 30},
		{"Asep", 35},
		{"Fuad", 23},
	}

	sort.Sort(UserSlice(users))

	fmt.Println(users)
}
