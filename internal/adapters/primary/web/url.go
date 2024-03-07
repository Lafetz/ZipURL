package web

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lafetz/url-shortner/internal/core/domain"
)

type createUrlBody struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	ShortUrl    string
	OriginalUrl string
}
type deleteUrl struct {
	Id uuid.UUID
}

func (a *App) createUrl(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		return
	}

	var ginUrl createUrlBody
	if err := c.ShouldBind(&ginUrl); err != nil {
		c.JSON(403, gin.H{
			"Errors": ValidateModel(err),
		})
		return
	}
	domainUrl := domain.NewUrl(user.(userToken).Id, ginUrl.OriginalUrl)
	_, err := a.urlService.AddUrl(domainUrl)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "url added",
		"url":     domainUrl,
	})

}

func (a *App) deleteUrl(c *gin.Context) {
	var url deleteUrl
	if err := c.ShouldBind(&url); err != nil {
		c.JSON(403, gin.H{
			"Errors": ValidateModel(err),
		})
	}
	err := a.urlService.DeleteUrl(url.Id)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"message": "url removed",
	})
}
func (a *App) getUrls(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		return
	}
	urls, err := a.urlService.GetUrls(user.(userToken).Id)
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"urls": urls,
	})
}
