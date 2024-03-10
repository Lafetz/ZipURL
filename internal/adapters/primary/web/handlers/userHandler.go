package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	jwtauth "github.com/lafetz/url-shortner/internal/adapters/primary/web/jwt"
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
				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"Errors": ValidateModel(err),
				})
				return

			}
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Error processing request body",
			})
			return
		}

		hashPassword, err := hashPassword(ginUser.Password)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return

		}
		domainUser := domain.NewUser(ginUser.Username, ginUser.Email, hashPassword)

		_, err = userService.AddUser(domainUser)

		if err != nil {

			if errors.Is(err, services.ErrEmailUnique) || errors.Is(err, services.ErrUsernameUnique) {
				c.JSON(http.StatusBadRequest, gin.H{
					"Error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "success,sigin to continue",
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

				c.JSON(http.StatusUnprocessableEntity, gin.H{
					"Errors": ValidateModel(err),
				})
				return

			}
			c.JSON(http.StatusBadRequest, gin.H{
				"Error": "Error processing request body",
			})
			return
		}
		user, err := userService.GetUser(ginUser.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Incorrect username or password",
			})
			return
		}
		err = matchPassword(ginUser.Password, user.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Incorrect username or password",
			})
			return
		}

		token, err := jwtauth.CreateJwt(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}

		c.SetCookie("Authorization", token, 24*60*60, "/", "localhost", true, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})

	}
}
