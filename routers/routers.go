package routers

import (
	"github.com/abrarnaim015/KP-golang-tgs7/controllers"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	return e
}