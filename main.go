package main

import (
	"absensi-app/frontend"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name" xml:"name" query:"name"`
	Email string `json:"email" xml:"email" query:"email"`
}

func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
	// or
	// return c.XML(http.StatusCreated, u)
}

func main() {
	e := echo.New()

	// init frontend
	frontend.RegisterHandlers(e)
	api := e.Group("api")
	api.POST("/save", saveUser)
	api.GET("/users/:id", getUser)
	api.GET("/users/show", show)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	e.Logger.Fatal(e.Start(":1323"))
}
