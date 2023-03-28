package main

import (
	"challenge-04/database"
	"challenge-04/models"
	"fmt"
)

// type User struct {
// 	gorm.Model
// 	Name  string `json:"name"`
// 	Email string `json:"email"`
// }

func main() {
	database.StartDB()

	createBook()
}

func createBook() {
	db := database.GetDB()

	Book := models.Book{
		ID:          7,
		Title:       "The Alchemist",
		Author:      "Paulo Coelho",
		Description: "The Alchemist is a novel by Paulo Coelho",
	}

	fmt.Println(Book, db)

	err := db.Create(&Book).Error

	if err != nil {
		fmt.Println("Error creating user data: ", err)
		return
	}

	fmt.Println("User created successfully")
}
