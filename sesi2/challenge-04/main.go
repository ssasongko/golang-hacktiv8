package main

import (
	"challenge-04/database"
	"challenge-04/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.InitDatabase()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	bookHandler := handler.NewBookHandler(db)

	r.GET("/books", bookHandler.GetBooks)
	r.POST("/books", bookHandler.CreateBook)
	r.GET("/books/:id", bookHandler.GetBookByID)
	r.PATCH("/books/:id", bookHandler.UpdateBook)
	r.DELETE("/books/:id", bookHandler.DeleteBook)

	r.Run(":8080")
}
