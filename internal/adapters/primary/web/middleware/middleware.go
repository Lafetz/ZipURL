package middleware

import (
	"github.com/gin-gonic/gin"
	jwt_auth "github.com/lafetz/url-shortner/internal/adapters/primary/web/jwt"
)

func RequireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("Authorization")
		if err != nil {
			c.JSON(401, gin.H{
				"Error": "Unauthorized",
			})
			return
		}

		user, err := jwt_auth.PareseJwt(jwtToken)
		if err != nil {
			if err == jwt_auth.ErrInvalidToken {
				c.JSON(401, gin.H{
					"Error": "Unauthorized",
				})
				return
			}

			c.JSON(500, gin.H{"Error": "Internal server error"})
			return

		}
		c.Set("user", user.GetUserToken())

		c.Next()
	}
}
