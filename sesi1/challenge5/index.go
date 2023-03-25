package main

import (
	"fmt"
	"sync"
	"time"
)

type Data struct {
	words []string
}

func printData(data Data, index int) {
	fmt.Printf("%v %d\n", data.words, index)
}

func printDataWithMutex(data Data, index int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	fmt.Printf("%v %d\n", data.words, index)
	mutex.Unlock()
}

func printDataRandomly(data1 Data, data2 Data) {
	for i := 0; i < 4; i++ {
		go printData(data1, i+1)
		go printData(data2, i+1)
	}
}

func printDataNeatly(data1 Data, data2 Data, mutex *sync.Mutex, wg *sync.WaitGroup) {
	for i := 0; i < 4; i++ {
		wg.Add(2)
		go printDataWithMutex(data1, i+1, mutex, wg)
		go printDataWithMutex(data2, i+1, mutex, wg)
		wg.Wait()
	}
}

func main() {
	data1 := Data{words: []string{"bisa1", "bisa2", "bisa3"}}
	data2 := Data{words: []string{"coba1", "coba2", "coba3"}}

	var mutex = &sync.Mutex{}
	var wg = &sync.WaitGroup{}

	printDataRandomly(data1, data2)
	time.Sleep(2 * time.Second)

	fmt.Printf("---------------------------------------------\n")

	printDataNeatly(data1, data2, mutex, wg)
}
