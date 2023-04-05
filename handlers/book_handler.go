package handlers

import (
	"net/http"
	"strconv"

	"gits-assignment/models"

	"github.com/gin-gonic/gin"
)

func GetBooks(c *gin.Context) {
	var book models.Book
	data, err := book.ReadBooks()
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func AddNewBook(c *gin.Context) {
	var req struct {
		Title    string `json:"title" binding:"required"`
		Year     uint   `json:"year" binding:"required"`
		AuthorID uint   `json:"author_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var book models.Book
	err := book.CreateBook(req.Title, req.Year, req.AuthorID)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "book created successfully",
	})
}

func ModifyBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Title    string `json:"title"`
		Year     uint   `json:"year"`
		AuthorID uint   `json:"author_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var book models.Book
	book, err := book.FindBookByID(id)
	if err != nil {
		AbortBadRequest(c, err)
		return
	}

	book.Title = req.Title
	book.Year = req.Year
	book.AuthorID = req.AuthorID

	err = book.UpdateBook(book)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "book updated successfully",
	})
}

func RemoveBook(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var book models.Book
	book, err := book.FindBookByID(id)
	if err != nil {
		AbortBadRequest(c, err)
		return
	}

	err = book.DeleteBook(book)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "book removed successfully",
	})
}
