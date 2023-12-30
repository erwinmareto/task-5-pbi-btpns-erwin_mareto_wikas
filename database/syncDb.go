package database

import "github.com/erwinmareto/profile-api-go/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{}, &models.Photo{})
}
