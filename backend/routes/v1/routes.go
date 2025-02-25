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
	g.GET("/:id", c.GetUserById)
	g.POST("", c.AdminCreateUser)
	g.PUT("/:id", c.AdminUpdateUser)
	g.DELETE("/:id", c.DeleteUser)
}

func RegisterRoutes(api *echo.Group, s RoutingControllers) {
	v1Group := api.Group("/v1")

	userGroup := v1Group.Group("/user")
	userRoutes(userGroup, s.UserController)
}
