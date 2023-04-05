package handlers

import (
	"net/http"

	"gits-assignment/models"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var newUser models.User
	err := newUser.CreateUser(req.Email, req.Password)
	if err != nil {
		AbortInternalServer(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "user successfully created",
	})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		AbortBadRequest(c, err)
		return
	}

	var user models.User
	result, err := user.FindUserByEmail(req.Email)
	if err != nil {
		AbortBadRequest(c, err)
		return
	}

	pass := user.CheckPassword(req.Password, result.Password)
	if pass != true {
		AbortBadRequest(c, err)
		return
	}

	token, _ := user.GenerateJWT(req.Email)

	c.JSON(http.StatusOK, gin.H{
		"msg":   "login success",
		"token": token,
	})
}
