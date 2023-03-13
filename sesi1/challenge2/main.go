package main

import "fmt"

func main() {

	// basic for loops
	for i := 0; i <= 4; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}

	for j := 0; j <= 10; j++ {
		// if j == 5 then print out the character
		if j == 5 {
			b := []rune{'С', 'ツ', 'А', 'ツ', 'Ш', 'ツ', 'А', 'ツ', 'Р', 'ツ', 'В', 'ツ', 'О'}

			for i := range b {
				// _ = i
				if i%2 == 0 {
					fmt.Printf("Character %U '%c' starts at byte position %d\n", b[i], b[i], i)
				}
			}
			continue
		}

		// if j == 6 then skip the character
		if j == 6 {
			continue
		}

		fmt.Printf("Nilai j = %d\n", j)
	}
}
