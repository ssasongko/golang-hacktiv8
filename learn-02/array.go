package main

import "fmt"

func main() {

	// first declare array
	var numbers [4]int
	// then assign value
	numbers = [4]int{1, 2, 3, 4}

	// or declare and assign value
	var strings = [4]string{"satu", "dua", "tiga", "empat"}

	fmt.Printf("%v \n", numbers)
	fmt.Printf("%#v \n", strings)

	// modify value
	strings[0] = "ichi"

	fmt.Println()

	// for loop through element
	for i, v := range strings {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println()

	// multidimensional array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}
