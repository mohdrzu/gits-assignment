package handlers

import (
	"net/http"
	"strconv"

	"gits-assignment/models"

	"github.com/gin-gonic/gin"
)

func GetAllPublishers(c *gin.Context) {

	var pub models.Publisher
	data, err := pub.ReadPublishers()
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})

}

func AddNewPublisher(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Location string `json:"location" binding:"required"`
		BookID   uint   `json:"book_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var pub models.Publisher
	err := pub.CreatePublisher(req.Name, req.Location, req.BookID)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "publisher created successfully",
	})
}

func ModifyPublisher(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var req struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		BookID   uint   `json:"book_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var pub models.Publisher
	pub, err := pub.FindPublisherByID(id)
	if err != nil {
		AbortBadRequest(c, err)
		return
	}

	pub.Name = req.Name
	pub.Location = req.Location
	pub.BookID = req.BookID

	err = pub.UpdatePublisher(pub)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "publisher updated successfully",
	})
}

func RemovePublisher(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var pub models.Publisher
	pub, err := pub.FindPublisherByID(id)
	if err != nil {
		AbortBadRequest(c, err)
		return
	}

	err = pub.DeletePublisher(pub)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "book removed successfully",
	})
}
