package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channel
	c1 := make(chan int)

	go func(c chan int) {
		fmt.Println("func goroutine start sending data into channel")
		c <- 10
		fmt.Println("func goroutine finish sending data into channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main goroutine start receiving data from channel")
	d := <-c1
	fmt.Println("main goroutine received data", d)

	close(c1)
	time.Sleep(time.Second)
}
