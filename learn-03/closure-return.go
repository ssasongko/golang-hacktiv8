package main

import (
	"fmt"
	"strings"
)

func main() {
	var studentLists = []string{"Airell", "Zahra", "Aaqila", "Sarah"}

	var find = findStudent(studentLists)

	fmt.Println(find("aaqila"))
	fmt.Println(find("ariel"))
}

func findStudent(students []string) func(string) string {

	// closure as return value
	return func(s string) string {
		var student string
		var position int

		for i, v := range students {
			if strings.ToLower(v) == strings.ToLower(s) {
				student = v
				position = i
				break
			}
		}

		if student == "" {
			return fmt.Sprintf("%s doesn't exist !!!\n", s)
		}
		return fmt.Sprintf("We found %s at the position of %d", s, position+1)
	}
}
