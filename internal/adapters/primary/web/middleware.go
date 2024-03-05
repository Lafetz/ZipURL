package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userToken struct {
	Id       uuid.UUID
	Email    string
	Username string
}

func requireAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtToken, err := c.Cookie("authorization")
		if err != nil {
			c.String(http.StatusNotFound, "Cookie not found")
			return
		}
		user, err := pareseJwt(jwtToken)
		if err != nil {
			return
		}
		c.Set("user", user)

		c.Next()
	}
}
