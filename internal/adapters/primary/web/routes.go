package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lafetz/url-shortner/internal/adapters/primary/web/handlers"
	"github.com/lafetz/url-shortner/internal/adapters/primary/web/middleware"
)

func (a *App) initAppRoutes() {

	a.gin.POST("/signup", handlers.CreateUser(a.userService))
	a.gin.POST("/signin", handlers.Signin(a.userService))
	//
	a.gin.GET("/urls", middleware.RequireAuth(), handlers.GetUrls(a.urlService))
	a.gin.POST("/urls", middleware.RequireAuth(), handlers.CreateUrl(a.urlService))
	a.gin.DELETE("/urls/:shorturl", middleware.RequireAuth(), handlers.DeleteUrl(a.urlService))
	//
	a.gin.GET("/ping", func(c *gin.Context) {
		c.String(200, "woking...")
	})
}
