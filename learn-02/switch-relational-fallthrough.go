package main

import "fmt"

func main() {
	var score = 4

	switch {
	case score == 8:
		fmt.Println("Perfect")
	case score < 8 && score > 3:
		fmt.Println("Not Bad")
		fallthrough // will continue to the next case
	case score < 2:
		fmt.Println("Your score was 2 what is wrong with you?")
	default:
		fmt.Println("Study harder")
		fmt.Println("You need learn more")
	}
}
