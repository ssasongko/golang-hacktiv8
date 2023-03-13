package main

import "fmt"

func main() {
	var a int = 1

	// there's no while in go
	for {
		fmt.Printf("Perulangan ke-%d\n", a)
		a++
		if a == 11 {
			break
		}
	}
	fmt.Println(a)
}
