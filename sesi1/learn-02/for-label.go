package main

import "fmt"

func main() {

outerLoop:
	for a := 0; a < 3; a++ {
		fmt.Printf("Perulangan ke - %d\n", a+1)
		for b := 0; b < 3; b++ {
			if a == 2 {
				break outerLoop
			}
			fmt.Print(b, " ")
		}
		fmt.Println()
	}
}
