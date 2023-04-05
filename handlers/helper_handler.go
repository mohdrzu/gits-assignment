package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AbortBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"error": fmt.Sprintf(err.Error()),
	})
}

func AbortInternalServer(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
		"error": fmt.Sprintf(err.Error()),
	})
}
