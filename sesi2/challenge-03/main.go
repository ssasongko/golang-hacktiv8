package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "ssasongko"
	dbname   = "hacktiv8_books"
)

var (
	db  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/books", GetBooksHandler)
	r.GET("/books/:bookID", GetBookHandler)
	r.POST("/book", CreateBookHandler)
	r.PATCH("/book/:bookID", UpdateBookHandler)
	r.DELETE("/book/:bookID", DeleteBookHandler)

	var PORT = ":8080"
	r.Run(PORT)
}

type Book struct {
	ID          int
	Title       string
	Author      string
	Description string
}

func GetBooksHandler(ctx *gin.Context) {
	var results = []Book{}

	sqlStatement := `SELECT * FROM books`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var book = Book{}

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

		if err != nil {
			panic(err)
		}

		results = append(results, book)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": results,
	})
}

func GetBookHandler(ctx *gin.Context) {
	var book Book

	row := db.QueryRow("SELECT * FROM books WHERE id = $1", ctx.Param("bookID"))
	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"books": book,
	})
}

func CreateBookHandler(ctx *gin.Context) {
	var book Book

	err := ctx.BindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var id int
	err = db.QueryRow("INSERT INTO books (title, author, description) VALUES ($1, $2, $3) RETURNING id", book.Title, book.Author, book.Description).Scan(&id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create book"})
		return
	}

	book.ID = id
	ctx.JSON(http.StatusOK, gin.H{
		"books": book,
	})
}

func UpdateBookHandler(ctx *gin.Context) {
	bookID := ctx.Param("bookID")

	var book Book
	err := ctx.BindJSON(&book)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var existingBook Book
	err = db.QueryRow("SELECT id, title, author, description FROM books WHERE id = $1", bookID).Scan(&existingBook.ID, &existingBook.Title, &existingBook.Author, &existingBook.Description)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if book.Title != "" {
		existingBook.Title = book.Title
	}
	if book.Author != "" {
		existingBook.Author = book.Author
	}
	if book.Description != "" {
		existingBook.Description = book.Description
	}

	_, err = db.Exec("UPDATE books SET title=$1, author=$2, description=$3 WHERE id=$4", existingBook.Title, existingBook.Author, existingBook.Description, existingBook.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": existingBook,
	})
}

func DeleteBookHandler(ctx *gin.Context) {
	var book Book
	bookID := ctx.Param("bookID")

	row := db.QueryRow("SELECT * FROM books WHERE id = $1", ctx.Param("bookID"))
	err = row.Scan(&book.ID, &book.Title, &book.Author, &book.Description)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	result, err := db.Exec("DELETE FROM books WHERE id = $1", bookID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
