package web

import "github.com/gin-gonic/gin"

func (a *App) initAppRoutes() {

	a.gin.POST("/signup", createUser(a.userService))
	a.gin.POST("/signin", signin(a.userService))
	//
	a.gin.GET("/urls", requireAuth(), getUrls(a.urlService))
	a.gin.POST("/urls", requireAuth(), createUrl(a.urlService))
	a.gin.DELETE("/urls/:shorturl", requireAuth(), deleteUrl(a.urlService))
	//
	a.gin.GET("/ping", func(c *gin.Context) {
		c.String(200, "woking...")
	})
}
