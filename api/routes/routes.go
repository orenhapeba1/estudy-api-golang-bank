package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/orenhapeba1/estudy-api-golang-bank/controllers"
	"github.com/orenhapeba1/estudy-api-golang-bank/middleware"
	_ "github.com/orenhapeba1/estudy-api-golang-bank/middleware"
)

func ConfigRoutes(router *gin.Engine) {

	main := router.Group("api/v1")
	{
		main.POST("login", controllers.Login)
		main.POST("accounts", controllers.AccountInsert)

		accounts := main.Group("accounts", middleware.Auth())
		{
			accounts.GET("/", controllers.AccountView)
			accounts.PUT("/", controllers.AccountUpdate)
			accounts.PUT("/:account", controllers.AccountUpdate)
			accounts.DELETE("/:account", controllers.AccountDelete)

		}

	}

}
