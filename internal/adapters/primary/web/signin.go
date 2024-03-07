package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type userSignin struct {
	Username string `json:"username" binding:"required,max=150"`
	Password string `json:"password" binding:"required,min=8,max=500" `
}

func (a *App) signinUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		var ginUser userSignin
		if err := c.ShouldBindJSON(&ginUser); err != nil {
			c.JSON(403, gin.H{
				"Errors": ValidateModel(err),
			})
			return
		}
		user, err := a.userService.GetUser(ginUser.Username)
		if err != nil {
			return
		}
		err = matchPassword(ginUser.Password, user.Password)
		if err != nil {
			return
		} //

		token, err := createJwt(user)
		if err != nil {
			return
		}
		//
		c.SetCookie("authorization", token, 24*60*60, "/", "localhost", true, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})

	}
}
