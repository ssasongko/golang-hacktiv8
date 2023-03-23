package main

import (
	"fmt"
)

func main() {
	var evenNumbers = func(numbers ...int) []int {
		var result []int

		for _, number := range numbers {
			if number%2 == 0 {
				result = append(result, number)
			}
		}

		return result
	}

	var number = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(evenNumbers(number...))

}
