package main

import "fmt"

func main() {
	for a := 0; a <= 4; a++ {
		for b := a; b <= 4; b++ {
			fmt.Print(b, " ")
		}
		fmt.Println()
	}
}
