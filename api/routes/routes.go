package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/orenhapeba1/estudy-api-golang-bank/controllers"
	_ "github.com/orenhapeba1/estudy-api-golang-bank/middleware"
)

func ConfigRoutes(router *gin.Engine) {

	main := router.Group("api/v1")
	{
		main.POST("login", controllers.Login)
		main.GET("/", controllers.Login)
	}

}
