package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	BookID string `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Books = []Book{
	{BookID: "b1", Title: "The Little Book of Ikigai: The Secret Japanese Way to Live a Happy and Long Life", Author: "Ken Mogi", Desc: "The Little Book of Ikigai is a Japanese concept that means finding your purpose in life. It is a combination of what you love, what you are good at, what the world needs, and what you can be paid for. The Little Book of Ikigai is a Japanese concept that means finding your purpose in life. It is a combination of what you love, what you are good at, what the world needs, and what you can be paid for."},
	{BookID: "b2", Title: "The Psychology of Money", Author: "Morgan Housels", Desc: "The Psychology of Money is a book by Morgan Housels that explores the psychology of money and how it affects our lives. The book is divided into three parts: The Psychology of Money, The Psychology of Money, and The Psychology of Money. The Psychology of Money is a book by Morgan Housels that explores the psychology of money and how it affects our lives. The book is divided into three parts: The Psychology of Money, The Psychology of Money, and The Psychology of Money."},
}

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"books": Books,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook Book

	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newBook.BookID = fmt.Sprintf("b%d", len(Books)+1)
	Books = append(Books, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range Books {
		if book.BookID == bookID {
			condition = true
			Books[i] = updatedBook
			Books[i].BookID = bookID
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Book not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfuly updated", bookID),
	})
}

func GetBook(ctx *gin.Context) {
	BookID := ctx.Param("bookID")
	condition := false
	var bookData Book

	for i, book := range Books {
		if book.BookID == BookID {
			condition = true
			bookData = Books[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Book not found",
			"error_message": fmt.Sprintf("Book with id %v not found", BookID),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"book": bookData,
	})
}

func DeleteBook(ctx *gin.Context) {
	bookID := ctx.Param("bookID")
	condition := false
	var bookIndex int

	for i, book := range Books {
		if bookID == book.BookID {
			condition = true
			bookIndex = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Book not found",
			"error_message": fmt.Sprintf("Book with id %v not found", bookID),
		})

		return
	}

	copy(Books[bookIndex:], Books[bookIndex+1:])
	Books[len(Books)-1] = Book{}
	Books = Books[:len(Books)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Book with id %v has been successfuly deleted", bookID),
	})
}
