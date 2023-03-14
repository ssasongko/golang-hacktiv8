package main

import "fmt"

func main() {
	for a := 0; a <= 10; a++ {
		// continue digunakan untuk melanjutkan perulangan
		if a%2 == 1 {
			continue
		}

		// break digunakan untuk menghentikan perulangan
		if a > 8 {
			break
		}

		fmt.Println("Perulangan ke ", a)
	}
}
