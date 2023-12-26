package main

import (
	"github.com/erwinmareto/database"
	"github.com/erwinmareto/helpers"
	"github.com/erwinmareto/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	helpers.LoadEnv()
	database.ConnectToDb()
	database.SyncDb()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "EYYY",
		})
	})
	routes.UserRouter(r)
	routes.PhotoRouter(r)
	r.Run() // listen and serve on 0.0.0.0:8080
} 