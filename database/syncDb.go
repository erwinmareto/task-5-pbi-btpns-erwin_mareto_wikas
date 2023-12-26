package database

import "github.com/erwinmareto/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Photo{})
}
