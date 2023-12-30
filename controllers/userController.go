package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/erwinmareto/profile-api-go/app"
	"github.com/erwinmareto/profile-api-go/database"
	"github.com/erwinmareto/profile-api-go/helpers"
	"github.com/gin-gonic/gin"
)

type UserData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func Register(c *gin.Context) {
	var user app.User

	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to read body",
		})
		return
	}
	if _, err := govalidator.ValidateStruct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := database.DB.Where("email = ?", user.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Email already exists",
		})
		return
	}

	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to hash password",
		})
		return
	}

	userData := app.User{Username: user.Username, Email: user.Email, Password: hash}

	result := database.DB.Create(&userData)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to create user",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User Registered Successfully",
	})
}

func Login(c *gin.Context) {
	var user app.User
	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to read body",
		})
		return
	}

	if err := database.DB.First(&user, app.User{Email: body.Email}).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	isMatched := helpers.CheckPasswordHash(body.Password, user.Password)

	if !isMatched {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	tokenString := helpers.GenerateToken(user.ID, user.Email)

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, app.LoginResponse{
		Success: true,
		Message: "Successfully logged in",
		Token:   tokenString,
		Data:    UserData{user.ID, user.Username, user.Email},
	})
}

func UpdateUser(c *gin.Context) {
	userId := c.Param("userId")
	idInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	var user app.User
	database.DB.First(&user, userId)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}
	currentUser, _ := c.Get("user")
	currentUserId := currentUser.(app.User).ID

	if uint(idInt) != currentUserId {
		c.JSON(http.StatusUnauthorized, app.ErrorResponse{
			Success: false,
			Message: fmt.Sprintf("You are not authorized to update user %d", currentUserId),
		})
		return
	}

	if c.Bind(&user) != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to read body",
		})
		return
	}
	hash, err := helpers.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to hash password",
		})
		return
	}

	user.Password = hash
	database.DB.Save(&user)
	c.JSON(http.StatusOK, app.SuccessResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    UserData{user.ID, user.Username, user.Email},
	})
}

func DeleteUser(c *gin.Context) {
	userId := c.Param("userId")
	idInt, err := strconv.Atoi(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	var user app.User
	database.DB.First(&user, userId)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "User not found",
		})
		return
	}

	currentUser, _ := c.Get("user")
	currentUserId := currentUser.(app.User).ID

	if uint(idInt) != currentUserId {
		c.JSON(http.StatusUnauthorized, app.ErrorResponse{
			Success: false,
			Message: "You are not authorized to delete this user",
		})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}
