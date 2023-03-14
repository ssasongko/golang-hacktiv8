package main

import "fmt"

func main() {
	var score = 5

	switch {
	case score == 3:
		fmt.Println("Perfect")
	case score < 8 && score > 3:
		fmt.Println("Not Bad")
	default:
		fmt.Println("Study harder")
		fmt.Println("You need learn more")
	}
}
