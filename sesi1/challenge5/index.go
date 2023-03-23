package main

import "fmt"

// Deklarasi Interface
type DataInterface1 interface {
	RetrieveData() []string
}

type Data1 struct {
	words []string
}

func (d Data1) RetrieveData() []string {
	return d.words
}

func main() {
	// deklarasi variabel c bertipe chan []string
	c := make(chan []string)

	data1 := Data1{words: []string{"bisa1", "bisa2", "bisa3"}}
	data2 := Data1{words: []string{"coba1", "coba2", "coba3"}}

	for i := 0; i < 4; i++ {
		go sendAndRetrieveData(data1, i+1, c)
	}

	for i := 0; i < 4; i++ {
		go sendAndRetrieveData(data2, i+1, c)
	}

	for i := 0; i < 8; i++ {
		data := <-c
		fmt.Println(data)
	}
}

func sendAndRetrieveData(dataRetriever DataInterface1, index int, c chan []string) {
	c <- dataRetriever.RetrieveData()
}
