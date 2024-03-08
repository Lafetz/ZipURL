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
	a.initAppRoutes()
	return a
}
func (a *App) Run() error {
	return a.gin.Run()
}
