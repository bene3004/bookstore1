package main

import (
	"awesomeProject1/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Book routes
	router.GET("/books", handlers.GetBooks)
	router.POST("/books", handlers.CreateBook)
	router.GET("/books/:id", handlers.GetBookByID)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	// Author routes
	router.GET("/authors", handlers.GetAuthors)
	router.POST("/authors", handlers.CreateAuthor)

	// Category routes
	router.GET("/categories", handlers.GetCategories)
	router.POST("/categories", handlers.CreateCategory)

	router.Run(":8080")
}
