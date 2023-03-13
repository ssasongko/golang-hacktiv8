package main

import (
	"fmt"
	"strings"
)

func main() {
	var text string = "Selamat Malam"

	pecahKata(text)
}

func pecahKata(text string) {
	var objectKata map[string]int
	objectKata = map[string]int{}

	for _, item := range text {
		var textLowerCase = strings.ToLower(string(item))

		fmt.Printf("%s \n", textLowerCase)

		var value, isExist = objectKata[textLowerCase]
		_ = value
		if !isExist {
			objectKata[textLowerCase] = 1
			continue
		}
		objectKata[textLowerCase] += 1
	}
	fmt.Println(objectKata)
}
