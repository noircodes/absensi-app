package routes

import (
	"absensi-app/backend/controllers"
	v1 "absensi-app/backend/routes/v1"
	"absensi-app/frontend"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routing(dblite *gorm.DB) *echo.Echo {
	e := echo.New()

	// init frontend
	frontend.RegisterHandlers(e)

	api := e.Group("/api")
	controllers := v1.RoutingControllers{
		UserController: controllers.UserController{},
	}

	v1.RegisterRoutes(api, controllers)

	return e
}
