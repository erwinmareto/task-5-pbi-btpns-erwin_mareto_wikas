package main

import (
	"github.com/erwinmareto/profile-api-go/database"
	"github.com/erwinmareto/profile-api-go/helpers"
	"github.com/erwinmareto/profile-api-go/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnv()
	database.ConnectToDb()
	database.SyncDb()
}

func main() {
	r := gin.Default()
	routes.UserRouter(r)
	routes.PhotoRouter(r)
	r.Run()
} 