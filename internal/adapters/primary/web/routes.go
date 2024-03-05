package web

import "github.com/gin-gonic/gin"

func (a *App) initAppRoutes() {

	a.gin.POST("/signup", a.createUser())
	a.gin.POST("/signin", a.signinUser())
	//
	a.gin.GET("/urls", a.getUrls)
	a.gin.POST("/urls", a.createUrl)
	a.gin.DELETE("/urls/:id", a.deleteUrl)
	//
	a.gin.GET("/ping", requireAuth(), func(c *gin.Context) {
		c.String(200, "yup working")
	})
}
