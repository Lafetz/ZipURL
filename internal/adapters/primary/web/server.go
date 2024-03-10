package web

import (
	"github.com/gin-gonic/gin"
	"github.com/lafetz/url-shortner/internal/core/services"
)

type App struct {
	userService services.UserServiceApi
	urlService  services.UrlServiceApi
	gin         *gin.Engine
	port        int
}

func NewApp(userService services.UserServiceApi, urlService services.UrlServiceApi) *App {
	a := &App{
		gin:         gin.Default(),
		urlService:  urlService,
		userService: userService,
		port:        8000,
	}
	a.gin.Use(corsMiddleware())
	a.initAppRoutes()

	return a
}
func (a *App) Run() error {
	return a.gin.Run()
}
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
