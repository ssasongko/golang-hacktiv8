package main

import "fmt"

func main() {
	c := make(chan string)

	students := []string{"Airell", "Nanda", "Mailo"}

	for _, v := range students {
		go introduce(v, c)
	}

	for i := 1; i <= len(students); i++ {
		print(c)
	}

	close(c)
}

// c <-- chan string is a receive-only channel
func print(c <-chan string) {
	fmt.Println(<-c)
}

// c chan<- string is a send-only channel
func introduce(student string, c chan<- string) {
	result := fmt.Sprintf("Hello, my name is %s", student)

	c <- result
}
