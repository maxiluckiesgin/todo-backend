package route

import (
	"todo-backend/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func Route(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())


	// enable CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		ExposeHeaders: []string{echo.HeaderContentDisposition},
	}))

	e.GET("/", controller.GetRoot)

}
