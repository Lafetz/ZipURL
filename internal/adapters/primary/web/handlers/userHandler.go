package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	jwt_auth "github.com/lafetz/url-shortner/internal/adapters/primary/web/jwt"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/services"
)

type UserCreateBody struct {
	Id       uuid.UUID `json:"id"  `
	Username string    `json:"username" binding:"required,max=150"`
	Email    string    `json:"email" binding:"required,email" `
	Password string    `json:"password" binding:"required,min=8,max=50" `
}

func CreateUser(userService services.UserServiceApi) gin.HandlerFunc {

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

		_, err = userService.AddUser(domainUser)

		if err != nil {
			fmt.Println(err)
			c.JSON(500, gin.H{
				"Error": "internal server Error",
			})
			return
		}

		c.JSON(201, gin.H{
			"message": "success",
			// "user":    user,
		})
	}

}

type userSignin struct {
	Username string `json:"username" binding:"required,max=150"`
	Password string `json:"password" binding:"required,min=8,max=500" `
}

func Signin(userService services.UserServiceApi) gin.HandlerFunc {
	return func(c *gin.Context) {

		var ginUser userSignin
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
		user, err := userService.GetUser(ginUser.Username)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Incorrect username or password",
			})
			return
		}
		err = matchPassword(ginUser.Password, user.Password)
		if err != nil {
			c.JSON(401, gin.H{
				"message": "Incorrect username or password",
			})
			return
		}

		token, err := jwt_auth.CreateJwt(user)
		if err != nil {
			c.Status(500)

			return
		}

		c.SetCookie("Authorization", token, 24*60*60, "/", "localhost", true, true)
		c.JSON(200, gin.H{
			"message": "success",
		})

	}
}
