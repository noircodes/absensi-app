package main

import (
	"absensi-app/backend/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s port=%s sslmode=disable TimeZone=%s dbname=%s",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASS"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_TIME"),
		os.Getenv("POSTGRES_DBNAME"),
	)
	fmt.Print(dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	e := routes.Routing(db)
	err = e.Start(fmt.Sprintf(":%d", 8080))
	if err != nil {
		return
	}
}
