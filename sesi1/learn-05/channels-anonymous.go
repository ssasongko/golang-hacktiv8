package main

import (
	"fmt"
)

func main() {
	// channels is a data structure that can be used to communicate between goroutines
	c := make(chan string)

	students := []string{"Airell", "Nanda", "Mailo"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hello, my name is %s", student)
			c <- result
		}(v)
	}

	for i := 1; i <= len(students); i++ {
		print(c)
	}

	close(c)
}

func print(c chan string) {
	fmt.Println(<-c)
}
