package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/erwinmareto/controllers"
)

func PhotoRouter(r *gin.Engine) {
	r.GET("/photos", controllers.GetPhotos)
	r.GET("/photos/:id", controllers.GetPhotoById)
	r.POST("/photos", controllers.CreatePhoto)
	r.PUT("/photos/:id", controllers.UpdatePhoto)
	r.DELETE("/photos/:id", controllers.DeletePhoto)
}