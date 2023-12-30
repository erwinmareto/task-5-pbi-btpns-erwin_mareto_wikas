package routes

import (
	"github.com/erwinmareto/profile-api-go/controllers"
	"github.com/erwinmareto/profile-api-go/middlewares"
	"github.com/gin-gonic/gin"
)

func PhotoRouter(r *gin.Engine) {
	r.Use(middlewares.RequireAuth)
	r.Use(middlewares.CheckAuthentication)
	r.GET("/photos", controllers.GetPhotos)
	r.GET("/photos/:id", controllers.GetPhotoById)
	r.POST("/photos", controllers.CreatePhoto)
	r.PUT("/photos/:id", controllers.UpdatePhoto)
	r.DELETE("/photos/:id", controllers.DeletePhoto)
}