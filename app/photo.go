package app

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	Title string `json:"title"`
	Caption string `json:"caption"`
	PhotoUrl string `json:"photo_url"`
	UserId int `json:"user_id"`
}