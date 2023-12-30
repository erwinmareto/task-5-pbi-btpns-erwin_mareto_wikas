package routes

import (
	"github.com/erwinmareto/profile-api-go/controllers"
	"github.com/erwinmareto/profile-api-go/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.PUT("/users/:userId", middlewares.RequireAuth, controllers.UpdateUser)
	r.DELETE("/users/:userId", middlewares.RequireAuth, controllers.DeleteUser)
}