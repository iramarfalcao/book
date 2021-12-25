package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/iramarfalcao/book/database"
	"github.com/iramarfalcao/book/models"
)

func GetBookById(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be Integer",
		})

		return
	}

	var book models.Book

	db := database.GetDatabase()
	err = db.First(&book, id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find book " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func CreateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func ListBooks(c *gin.Context) {
	db := database.GetDatabase()

	var books []models.Book

	err := db.Find(&books).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, books)
}

func UpdateBook(c *gin.Context) {
	db := database.GetDatabase()

	var book models.Book

	err := c.ShouldBindJSON(&book)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&book).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create book: " + err.Error(),
		})

		return
	}

	c.JSON(200, book)
}

func DeleteBook(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be Integer",
		})

		return
	}

	db := database.GetDatabase()
	err = db.Delete(&models.Book{}, id).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete book " + err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"message": "book successfully deleted",
	})
}
