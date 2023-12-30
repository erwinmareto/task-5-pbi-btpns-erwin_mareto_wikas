package models

import "github.com/erwinmareto/profile-api-go/app"

type User struct {
	app.User
	Photos []Photo `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
