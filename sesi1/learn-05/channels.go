package main

import (
	"fmt"
)

func main() {
	// channels is a data structure that can be used to communicate between goroutines
	c := make(chan string)

	go introduce("Airell", c)
	go introduce("Nanda", c)
	go introduce("Mailo", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)
}

func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hello, my name is %s", student)

	c <- result
}
