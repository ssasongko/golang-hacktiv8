package main

import "fmt"

func main() {
	var currentYear = 2023

	if age := currentYear - 2001; age < 17 {
		fmt.Println("You are not old enough to drive")
	} else {
		fmt.Println("You are old enough to drive")
	}
}
