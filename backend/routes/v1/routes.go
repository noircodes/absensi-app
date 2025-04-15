package v1

import (
	"absensi-app/backend/controllers"

	"github.com/labstack/echo/v4"
)

type RoutingControllers struct {
	UserController controllers.UserController
}

func userRoutes(g *echo.Group, c controllers.UserController) {
	g.GET("/all", c.AdminGetAllUser)
	g.GET("/get/:id", c.GetUserById)
	g.POST("/create", c.AdminCreateUser)
	g.PUT("/update/:id", c.AdminUpdateUser)
	g.DELETE("/delete/:id", c.DeleteUser)
}

func RegisterRoutes(api *echo.Group, s RoutingControllers) {
	v1Group := api.Group("/v1")

	userGroup := v1Group.Group("/user")
	userRoutes(userGroup, s.UserController)
}
