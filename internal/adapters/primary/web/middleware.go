package web

import (
	"github.com/gin-gonic/gin"
)

func requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(401, gin.H{
				"Error": "Unauthorized",
			})
			return
		}

		user, err := pareseJwt(jwtToken)
		if err != nil {
			if err == ErrInvalidToken {
				c.JSON(401, gin.H{
					"Error": "Unauthorized",
				})
				return
			}

			c.JSON(500, gin.H{"Error": "Internal server error"})
			return

		}
		c.Set("user", user)

		c.Next()
	}
}
