package handlers

import (
	"net/http"
	"strconv"

	"gits-assignment/models"

	"github.com/gin-gonic/gin"
)

func GetAllAuthors(c *gin.Context) {

	var author models.Author
	data, err := author.ReadAuthors()
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func AddNewAuthor(c *gin.Context) {
	var req struct {
		Name  string `json:"name" binding:"required"`
		Email string `json:"email" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var author models.Author
	err := author.CreateAuthor(req.Name, req.Email)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "author created successfully",
	})
}

func ModifyAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var author models.Author
	author, err := author.FindAuthorByID(id)
	if err != nil {
		AbortBadRequest(c, err)
		return
	}

	author.Name = req.Name
	author.Email = req.Email

	err = author.UpdateAuthor(author)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "author updated successfully",
	})
}

func RemoveAuthor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var author models.Author
	author, err := author.FindAuthorByID(id)
	if err != nil {
		AbortBadRequest(c, err)
		return
	}

	err = author.DeleteAuthor(author)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "author removed successfully",
	})
}
