package router

import (
	"api_test/controllers"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	fmt.Println("Welcome to the App Test")
	e := echo.New()

	// create groups

	// adminGroup := e.Group("/admin")
	// cookieGroup := e.Group("/cookie")
	// jwtGroup := e.Group("/jwt")

	// set all middlewares

	// middlewares.SetMainMiddlewares(e)
	// middlewares.SetAdminMiddlewares(adminGroup)
	// middlewares.SetCookieMiddlewares(cookieGroup)
	// middlewares.SetJwtMiddlewares(jwtGroup)

	// set main routes

	// api.MainGroup(e)

	// set group routes

	// api.AdminGroup(adminGroup)
	// api.CookieGroup(cookieGroup)
	// api.JwtGroup(jwtGroup)

	e.POST("/getuser", controllers.GetUser)
	e.POST("/login", controllers.Login)
	e.POST("/RegisterAccount", controllers.RegisterAccount)

	return e
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}
