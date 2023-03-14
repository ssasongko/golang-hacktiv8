package main

import "fmt"

func main() {
	var a int = 1

	// there's no while in go
	for a <= 10 {
		fmt.Printf("Perulangan ke-%d\n", a)
		a++
	}
	fmt.Println(a)
}
