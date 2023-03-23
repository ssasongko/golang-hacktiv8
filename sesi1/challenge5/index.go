package main

import (
	"fmt"
)

type Data1 struct {
	Value []string
}

type DataRetriever interface {
	RetrieveData() []string
}

func (d Data1) RetrieveData() []string {
	return d.Value
}

func main() {
	c := make(chan []string)

	data1 := Data1{
		Value: []string{"bisa1", "bisa2", "bisa3"},
	}

	data2 := Data1{
		Value: []string{"coba1", "coba2", "coba3"},
	}

	for i := 0; i < 4; i++ {
		go myFunction(data1, c)
	}

	for i := 0; i < 4; i++ {
		go myFunction(data2, c)
	}

	for i := 0; i < 8; i++ {
		data := <-c
		fmt.Println(data)
	}

	close(c)
}

func myFunction(dataRetriever DataRetriever, c chan []string) {
	c <- dataRetriever.RetrieveData()
}
