package app

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string `json:"username" valid:"required~Username is required" gorm:"not null"`
	Email     string `json:"email" valid:"required~Email is required, email~Invalid email address" gorm:"unique"`
	Password  string `json:"password" valid:"required~Password is required, minstringlength(6)~Password has to have a minimum length of 6 characters"`
}