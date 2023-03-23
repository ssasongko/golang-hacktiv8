package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// first function (void)
	greet("Airell", 18)

	fmt.Println(strings.Repeat("-", 20))

	// second double return function
	var diameter float64 = 15
	var area, circumference = calcuate(diameter)

	fmt.Println("Area: ", area)
	fmt.Println("Circumference: ", circumference)

	fmt.Println(strings.Repeat("-", 20))

	// third double return function with predefined return value
	// double return function
	var diameter2 float64 = 20
	var area2, circumference2 = calcuate(diameter2)

	fmt.Println("Area: ", area2)
	fmt.Println("Circumference: ", circumference2)
}

// # 1
func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old\n", name, age)
}

// # 2
func calcuate(d float64) (float64, float64) {
	// menghitung luas
	var area float64 = math.Pi * math.Pow(d/2, 2)

	// menghitung keliling
	var circumference float64 = math.Pi * d

	return area, circumference
}

// # 3 : Predefined function
func calcuate2(d float64) (area2 float64, circumference2 float64) {
	// menghitung luas
	area2 = math.Pi * math.Pow(d/2, 2)

	// menghitung keliling
	circumference2 = math.Pi * d

	return
}
