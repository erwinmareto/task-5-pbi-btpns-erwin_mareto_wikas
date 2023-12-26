package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/erwinmareto/controllers"
)

func UserRouter(r *gin.Engine) {
	r.GET("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)
}