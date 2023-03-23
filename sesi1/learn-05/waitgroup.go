package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fruits := []string{"apple", "manggo", "durian", "rambutan"}

	// create a WaitGroup
	var wg sync.WaitGroup

	for index, fruit := range fruits {
		wg.Add(1) // add 1 to the WaitGroup counter
		go printFruit(index, fruit, &wg)
	}

	fmt.Println("No. of Goroutines", runtime.NumGoroutine())

	wg.Wait()
}

func printFruit(index int, fruit string, wg *sync.WaitGroup) {
	fmt.Printf("index => %d, fruit => %s\n", index, fruit)
	wg.Done()
}
