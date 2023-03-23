package main

import (
	"fmt"
	"time"
)

func main() {
	// unbuffered channel
	c1 := make(chan int, 3)

	go func(c chan int) {
		for i := 0; i <= 5; i++ {
			fmt.Println("func goroutine start sending data into channel")
			c <- i
			fmt.Println("func goroutine finish sending data into channel")
		}

		close(c)
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	for v := range c1 {
		fmt.Println("main goroutine received value from channel", v)
	}
}
