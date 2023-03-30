package handler

import (
	"net/http"

	"challenge-04/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookHandler struct {
	db *gorm.DB
}

func NewBookHandler(db *gorm.DB) *BookHandler {
	return &BookHandler{db}
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	var books []models.Book
	if err := h.db.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Books fetched successfully",
		"data":    books,
	})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := h.db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Book fetched successfully",
		"data":    book,
	})
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if err := h.db.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Book created successfully",
		"data":    book,
	})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := h.db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if err := h.db.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Book updated successfully",
		"data":    book,
	})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := h.db.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	if err := h.db.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  true,
		"message": "Book deleted successfully",
	})
}
