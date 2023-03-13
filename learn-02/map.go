package main

import (
	"fmt"
	"strings"
)

func main() {

	// declare 1
	var person map[string]string
	person = map[string]string{}

	person["name"] = "John"
	person["title"] = "Programmer"
	person["age"] = "25"

	fmt.Printf("%v", person)

	fmt.Println(strings.Repeat("=", 50))

	// declare 2
	var person2 = map[string]string{
		"name":  "John",
		"title": "Programmer",
		"age":   "25",
	}
	fmt.Printf("%v\n", person2)

	fmt.Println(strings.Repeat("=", 50))

	// print out
	for key, value := range person2 {
		fmt.Printf("%s: %s \n", key, value)
	}

	fmt.Println(strings.Repeat("=", 50))

	// detect key
	var value, isExist = person2["age"]
	if isExist {
		fmt.Printf("Key ada ges: %s \n", value)
	} else {
		fmt.Println("Key tidak ada")
	}

	// delete
	delete(person2, "age")
	for key, value := range person2 {
		fmt.Printf("%s: %s \n", key, value)
	}
	fmt.Println(strings.Repeat("=", 50))

	// detect key
	var value2, isExist2 = person2["age"]
	if isExist2 {
		fmt.Printf("Key ada ges: %s \n", value2)
	} else {
		fmt.Println("Key tidak ada")
	}

	fmt.Println(strings.Repeat("=", 50))

	var persons = []map[string]string{
		{"name": "John", "title": "Programmer", "age": "25"},
		{"name": "Maya", "title": "QA Tester", "age": "19"},
		{"name": "Maya", "title": "QA Tester", "age": "19"},
	}
	// map combining slice
	for i, item := range persons {
		fmt.Printf("Index: %d, Name: %s, Age: %s\n", i, item["name"], item["title"])
	}
}
