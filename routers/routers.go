package routers

import (
	auth_controllers "bookstore/controllers/auth"
	categories_controllers "bookstore/controllers/categories"
	"bookstore/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	users := api.Group("/users")
	categories := api.Group("/categories", middleware.AuthValidation())

	{
		users.POST("/register", auth_controllers.Register)
		users.POST("/login", auth_controllers.Login)

		categories.POST("/", categories_controllers.CreateCategories)
		categories.GET("/", categories_controllers.GetAllCategories)
		categories.GET("/:id", categories_controllers.GetCategoriesById)
		categories.PATCH("/:id", categories_controllers.UpdateCategories)
		categories.DELETE("/:id", categories_controllers.DeleteCategoryById)
	}

	return router
}
