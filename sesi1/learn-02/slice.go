package main

import "fmt"

func main() {
	// berbeda dengan array, Slice tidak memiliki sifat fixed-length
	// deklarasi slice
	var fruits = []string{"apple", "grape", "banana", "melon"}

	// make function
	var fruits2 = make([]string, 3)

	_, _ = fruits, fruits2

	fmt.Printf("%#v\n", fruits)
	fmt.Printf("%#v\n", fruits2)

	fmt.Println()

	// append
	fruits2 = append(fruits2, "orange", "papaya", "mango")
	fmt.Printf("%#v\n", fruits2)

	fmt.Println()

	// ellipsis
	var fruits3 = append(fruits, fruits2...)
	fmt.Printf("%#v\n", fruits3)

	fmt.Println()

	// copy
	var seaAnimals = []string{"shark", "salmon", "whale", "octopus"}
	var animals = []string{"lion", "monkey", "dog", "cat"}

	// copy function
	// 1st parameter is destination
	// 2nd parameter is source
	var animalsCopy = copy(seaAnimals, animals)
	fmt.Printf("%#v\n", seaAnimals)
	fmt.Printf("%#v\n", animals)
	fmt.Printf("%#v\n", animalsCopy)

	fmt.Println()

	// slicing
	var fruits4 = []string{"apple", "grape", "banana", "melon"}
	var newFruits = fruits4[1:3]
	fmt.Printf("%#v\n", newFruits)
}
