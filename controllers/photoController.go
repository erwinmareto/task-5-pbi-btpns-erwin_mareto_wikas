package controllers

import (
	"net/http"
	"strconv"

	"github.com/erwinmareto/profile-api-go/app"
	"github.com/erwinmareto/profile-api-go/database"
	"github.com/gin-gonic/gin"
)

type PhotoData struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId   int    `json:"user_id"`
}

func GetPhotos(c *gin.Context) {
	var photos []app.Photo
	database.DB.Find(&photos)
	c.JSON(http.StatusOK, gin.H{
		"data": photos,
	})
}

func GetPhotoById(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to read body",
		})
		return
	}
	var photo app.Photo
	database.DB.First(&photo, id)

	if photo.ID == 0 {
		c.JSON(http.StatusNotFound, app.ErrorResponse{
			Success: false,
			Message: "Photo not found",
		})
		return
	}
	c.JSON(http.StatusOK, app.SuccessResponse{
		Success: true,
		Message: "Photo retrieved successfully",
		Data: PhotoData{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserId:   photo.UserId,
		},
	})
}

func CreatePhoto(c *gin.Context) {
	var photo app.Photo
	if c.Bind(&photo) != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to read body",
		})
		return
	}

	user, _ := c.Get("user")
	userId := user.(app.User).ID

	photo.UserId = int(userId)
	result := database.DB.Create(&photo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to create photo",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo uploaded successfully",
	})
}

func UpdatePhoto(c *gin.Context) {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	var photoData app.Photo
		user, _ := c.Get("user")
		userId := user.(app.User).ID
	database.DB.Where("ID = ?", idInt).First(&photoData)

	if photoData.UserId != int(userId) {
		c.JSON(http.StatusUnauthorized, app.ErrorResponse{
			Success: false,
			Message: "You are not authorized to update this photo",
		})
		return
	}

	if c.Bind(&photoData) != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to read body",
		})
		return
	}
	photo := app.Photo{UserId: int(userId), Title: photoData.Title, Caption: photoData.Caption, PhotoUrl: photoData.PhotoUrl}
	photo.ID = uint(idInt)
	result := database.DB.Save(&photo)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to update photo",
		})
		return
	}

	c.JSON(http.StatusOK, app.SuccessResponse{
		Success: true,
		Message: "Photo Updated Successfully",
		Data: PhotoData{
			ID:       photo.ID,
			Title:    photo.Title,
			Caption:  photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserId:   photo.UserId,
		},
	})
}

func DeletePhoto(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Invalid ID",
		})
		return
	}

	var photoData app.Photo

	user, _ := c.Get("user")
	userId := user.(app.User).ID
	database.DB.Where("ID = ?", idInt).First(&photoData)

	if photoData.UserId != int(userId) {
		c.JSON(http.StatusUnauthorized, app.ErrorResponse{
			Success: false,
			Message: "You are not authorized to delete this photo",
		})
		return
	}

	if err := database.DB.Delete(&app.Photo{}, idInt).Error; err != nil {
		c.JSON(http.StatusBadRequest, app.ErrorResponse{
			Success: false,
			Message: "Failed to delete photo",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Photo Deleted Successfully.",
	})
}
