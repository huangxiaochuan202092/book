package handlers

import (
	"book-management-api/database"
	"book-management-api/models"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&book)
	c.JSON(200, gin.H{ //返回json
		"book": book,
	})
}

func GetBooks(c *gin.Context) {
	var books []models.Book
	database.DB.Find(&books)
	c.JSON(200, books)

}

func GetBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(200, book)
}





func UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	database.DB.Save(&book)
	c.JSON(200, book)

}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")
	var book models.Book
	if err := database.DB.Where("id = ?", id).First(&book).Error; err != nil {
		c.JSON(404, gin.H{"error": "Book not found"})
		return
	}
	database.DB.Delete(&book)
	c.JSON(200, gin.H{ //返回json
		"message": "Book deleted",
	})
}
