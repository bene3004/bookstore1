package handlers

import (
	"net/http"
	"strconv"

	"awesomeProject1/models"

	"github.com/gin-gonic/gin"
)

var books = []models.Book{}
var nextBookID = 1

// List books with pagination and filters
func GetAllBooks(c *gin.Context) {
	categoryFilter := c.Query("category")
	authorFilter := c.Query("author")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "5"))

	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 5
	}

	var filtered []models.Book

	// filtering based on category and author
	for _, b := range books {
		if (categoryFilter == "" || strconv.Itoa(b.CategoryID) == categoryFilter) && (authorFilter == "" || strconv.Itoa(b.AuthorID) == authorFilter) {
			filtered = append(filtered, b)
		}
	}

	// pagination
	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if startIndex > len(filtered) {
		c.JSON(http.StatusOK, []models.Book{})
		return
	}

	if endIndex > len(filtered) {
		endIndex = len(filtered)
	}

	c.JSON(http.StatusOK, filtered[startIndex:endIndex])
}

// Add a new book
func CreateBook(c *gin.Context) {
	var newBook models.Book
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newBook.ID = nextBookID
	nextBookID++
	books = append(books, newBook)
	c.JSON(http.StatusCreated, newBook)
}

// Get a book by ID
func GetBookByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for _, book := range books {
		if book.ID == id {
			c.JSON(http.StatusOK, book)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// Update a book
func UpdateBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var updatedBook models.Book
	if err := c.ShouldBindJSON(&updatedBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, book := range books {
		if book.ID == id {
			updatedBook.ID = id
			books[i] = updatedBook
			c.JSON(http.StatusOK, updatedBook)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}

// Delete a book
func DeleteBook(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
}
