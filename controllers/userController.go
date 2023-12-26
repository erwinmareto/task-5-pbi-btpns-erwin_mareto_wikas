package controllers

import (
	"net/http"

	"github.com/erwinmareto/database"
	"github.com/erwinmareto/helpers"
	"github.com/erwinmareto/models"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var body struct {
		Username string
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := helpers.HashPassword(body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}

	user := models.User{Username: body.Username, Email: body.Email, Password: hash}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all good REGISTERED UP",
	})
}

func Login(c *gin.Context) {
	var body struct {
		Username string
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	var user models.User
	database.DB.First(&user, models.User{Email: body.Email})

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err := helpers.CheckPasswordHash(body.Password, user.Password)

	if err {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "all good logged in",
		"data":    user,
	})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	var body struct {
		Username string
		Email    string
		Password string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	var user models.User
	database.DB.First(&user, userId)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}
	user.Username = body.Username
	user.Email = body.Email
	user.Password = body.Password
	database.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"data":    user,
	}) 
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	var user models.User
	database.DB.First(&user, userId)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User not found",
		})
		return
	}
	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
