package web

func (a *App) initAppRoutes() {

	a.gin.POST("/signup", a.createUser())
	a.gin.POST("/signin", a.signinUser())
	// a.gin.POST("/signin", a.SignInUser(a.userService))
}
