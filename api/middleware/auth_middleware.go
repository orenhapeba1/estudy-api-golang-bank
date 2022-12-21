package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/orenhapeba1/estudy-api-golang-bank/services"
)

func Auth() gin.HandlerFunc {

	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")

		if header == "" {
			c.AbortWithStatus(406)
		}

		token := header[len(Bearer_schema):]
		if !services.NewJWTService().ValidateToken(token) {
			c.AbortWithStatus(401)
		}

	}
}
