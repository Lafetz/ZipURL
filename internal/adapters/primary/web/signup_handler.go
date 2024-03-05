package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
)

type UserCreateBody struct {
	Id       uuid.UUID `json:"id"  `
	Username string    `json:"username" binding:"required,max=150"`
	Email    string    `json:"email" binding:"required,email" `
	Password string    `json:"password" binding:"required,min=8,max=500" `
}

func (a *App) createUser() gin.HandlerFunc {

	return func(c *gin.Context) {
		var ginUser UserCreateBody
		if err := c.ShouldBindJSON(&ginUser); err != nil {
			c.JSON(403, gin.H{
				"Errors": ValidateModel(err),
			})
			return
		}

		hashPassword, err := hashPassword(ginUser.Password)

		if err != nil {
			c.JSON(500, gin.H{
				"Error": "internal server Error",
			})
			return

		}
		domainUser := domain.NewUser(ginUser.Username, ginUser.Email, hashPassword)

		user, err := a.userService.AddUser(domainUser)

		if err != nil {
			return
		}
		fmt.Printf("user is %s", user)
		c.JSON(200, gin.H{
			"message": "success",
			// "user":    user,
		})
	}

}
