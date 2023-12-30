package models

import 
	"github.com/erwinmareto/profile-api-go/app"

type Photo struct {
	app.Photo
	User User
}