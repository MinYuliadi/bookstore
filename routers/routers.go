package routers

import (
	controller "bookstore/controllers/auth"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	users := api.Group("/users")

	{
		users.POST("/register", controller.Register)
		users.POST("/login", controller.Login)
	}

	return router
}
