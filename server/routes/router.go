package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iramarfalcao/book/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		books := main.Group("books")
		{
			books.GET("/", controllers.GetBookById)
		}
	}

	return router
}
