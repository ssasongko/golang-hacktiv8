package main

import (
	"fmt"
)

func main() {
	// immediately invoked function expression
	// function yang langsung dijalankan

	var isPalindrome = func(value string) bool {
		var str string = ""

		for i := len(value) - 1; i >= 0; i-- {
			str += string(byte(value[i]))
		}

		return str == value
	}("katak")

	fmt.Println(isPalindrome)

}
