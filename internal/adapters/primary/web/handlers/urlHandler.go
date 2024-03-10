package handlers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	jsonformatter "github.com/lafetz/url-shortner/internal/adapters/primary/web/formatter"
	jwtauth "github.com/lafetz/url-shortner/internal/adapters/primary/web/jwt"
	"github.com/lafetz/url-shortner/internal/core/domain"
	"github.com/lafetz/url-shortner/internal/core/services"
)

type createUrlReq struct {
	OriginalUrl string `json:"originalUrl" binding:"required,min=5"`
}

func CreateUrl(urlService services.UrlServiceApi) gin.HandlerFunc {
	return func(c *gin.Context) {

		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Unauthorized",
			})
			return
		}
		var ginUrl createUrlReq
		if err := c.ShouldBind(&ginUrl); err != nil {
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

		id, err := uuid.Parse(user.(*jwtauth.UserToken).Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}
		domainUrl := domain.NewUrl(id, ginUrl.OriginalUrl)
		url, err := urlService.AddUrl(domainUrl)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}

		res := jsonformatter.NewUrlResp(url)

		c.JSON(http.StatusCreated, gin.H{
			"message": "url added",
			"url":     res,
		})

	}
}

func DeleteUrl(urlService services.UrlServiceApi) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Unauthorized",
			})
			return
		}
		shortUrl := c.Param("shorturl")
		userId, err := uuid.Parse(user.(*jwtauth.UserToken).Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}
		err = urlService.DeleteUrl(shortUrl, userId)
		if err != nil {
			if errors.Is(err, services.ErrUrlNotFound) {
				c.JSON(http.StatusNotFound, gin.H{
					"Error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "url removed",
		})
	}
}

func GetUrls(urlService services.UrlServiceApi) gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Unauthorized",
			})
			return
		}
		id, err := uuid.Parse(user.(*jwtauth.UserToken).Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}
		urls, err := urlService.GetUrls(id)
		if err != nil {
			if err != nil {
				if errors.Is(err, services.ErrUrlNotFound) {
					c.JSON(http.StatusNotFound, gin.H{
						"Error": err.Error(),
					})
					return
				}
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}

		res := jsonformatter.NewUrlsResp(urls)

		c.JSON(http.StatusOK, gin.H{
			"urls": res,
		})
	}
}
func Redirect(urlservice services.UrlServiceApi) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortUrl := c.Param("shorturl")

		url, err := urlservice.GetUrl(shortUrl)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{
				"Error": "internal server Error",
			})
			return
		}

		c.Redirect(http.StatusSeeOther, url.OriginalUrl)

	}
}
