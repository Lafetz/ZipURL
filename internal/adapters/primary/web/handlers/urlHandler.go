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

type createUrlReq struct {
	OriginalUrl string `json:"originalUrl" binding:"required,min=5" `
}

func CreateUrl(urlService services.UrlServiceApi) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(401, gin.H{
				"Error": "Unauthorized",
			})
		}
		var ginUrl createUrlReq
		if err := c.ShouldBind(&ginUrl); err != nil {
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

		id, err := uuid.Parse(user.(*jwt_auth.UserToken).Id)
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Internal Server Error",
			})
		}
		domainUrl := domain.NewUrl(id, ginUrl.OriginalUrl)
		_, err = urlService.AddUrl(domainUrl)
		if err != nil {
			fmt.Print(err)
			c.JSON(500, gin.H{
				"Error": "Internal Server Error",
			})
			return
		}
		c.JSON(201, gin.H{
			"message": "url added",
			"url":     domainUrl,
		})

	}
}

func DeleteUrl(urlService services.UrlServiceApi) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(401, gin.H{
				"Error": "Unauthorized",
			})
			return
		}
		id := c.Param("shorturl")
		userId, err := uuid.Parse(user.(*jwt_auth.UserToken).Id)
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Internal Server Error",
			})
		}
		err = urlService.DeleteUrl(id, userId)
		if err != nil {
			fmt.Print(err)
			c.JSON(500, gin.H{
				"Error": "Internal Server Error",
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "url removed",
		})
	}
}

func GetUrls(urlService services.UrlServiceApi) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			return
		}
		id, err := uuid.Parse(user.(*jwt_auth.UserToken).Id)
		if err != nil {
			c.JSON(500, gin.H{
				"Error": "Internal Server Error",
			})
		}
		urls, err := urlService.GetUrls(id)
		if err != nil {
			fmt.Print(err)
			c.JSON(500, gin.H{
				"Error": "Internal Server Error",
			})
			return
		}
		c.JSON(200, gin.H{
			"urls": urls,
		})
	}
}