package web

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/services"
)

type UserCreateBody struct {
	Id       uuid.UUID `json:"id"  `
	Username string    `json:"username" binding:"required,max=150"`
	Email    string    `json:"email" binding:"required,email" `
	Password string    `json:"password" binding:"required,min=8,max=50" `
}

// (a *App)
func createUser(userService services.UserServicePort) gin.HandlerFunc {

	return func(c *gin.Context) {
		var ginUser UserCreateBody
		if err := c.ShouldBindJSON(&ginUser); err != nil {
			_, ok := err.(validator.ValidationErrors)
			if ok {

				c.JSON(422, gin.H{
					"Errors": ValidateModel(err),
				})
				return

			}
			c.JSON(400, gin.H{
				"Error": "Error processing request body",
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

		user, err := userService.AddUser(domainUser)

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
