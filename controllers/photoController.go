package controllers

import (
	"net/http"
	"strconv"

	"github.com/erwinmareto/database"
	"github.com/erwinmareto/models"
	"github.com/gin-gonic/gin"
)

func GetPhotos(c *gin.Context) {
	var photos []models.Photo
	database.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{
		"data": photos,
	})
}

func GetPhotoById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	var photo models.Photo
	database.DB.First(&photo, id)

	if photo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Photo not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": photo,
	})
}

func CreatePhoto(c *gin.Context) {
	var body struct {
		Title string
		Caption string
		PhotoUrl string
		UserId int
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	photo := models.Photo{Title: body.Title, Caption: body.Caption, PhotoUrl: body.PhotoUrl, UserId: body.UserId}
	result := database.DB.Create(&photo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create photo",
		})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Photo uploaded yep.",
	})
}

func UpdatePhoto(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	var body struct {
		UserId int
		Title string
		Caption string
		PhotoUrl string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	photo := models.Photo{UserId: body.UserId, Title: body.Title, Caption: body.Caption, PhotoUrl: body.PhotoUrl}
	photo.ID = uint(idInt)
	result := database.DB.Save(&photo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update photo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo updated, nice.",
		"data": photo,
	})
}

func DeletePhoto(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid ID",
		})
		return
	}
	result := database.DB.Delete(&models.Photo{}, idInt)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to delete photo",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Photo deleted, yay.",
	})
}
