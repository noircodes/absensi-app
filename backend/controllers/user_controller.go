package controllers

import (
	"absensi-app/backend/models"
	"absensi-app/backend/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService services.UserService
}

func (h *UserController) GetAllUser(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func (h *UserController) GetUserById(c echo.Context) error {
	id := c.Param("id")

	user, err := h.userService.GetUserById(id)
	if err != nil {
		return c.String(http.StatusNotFound, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *UserController) AdminGetAllUser(c echo.Context) error {
	pageParam := c.QueryParam("page")
	limitParam := c.QueryParam("limit")
	name := c.QueryParam("name")
	role := c.QueryParam("role")
	email := c.QueryParam("email")
	// code := c.QueryParam("code")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 {
		limit = 10
	}

	users, total, _ := h.userService.GetAllUsers(page, limit, name, role, email)
	totalPages := (total + limit - 1) / limit
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":    "success",
		"total":      total,
		"page":       page,
		"limit":      limit,
		"totalPages": totalPages,
		"data":       users,
	})
}

func (h *UserController) AdminCreateUser(c echo.Context) error {
	var req models.CreateUserRequest
	fmt.Print(&req)
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	h.userService.CreateUser(&req)

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully created."})
}

func (h *UserController) AdminUpdateUser(c echo.Context) error {
	id := c.Param("id")

	var req models.UpdateUserRequest
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request"})
	}

	h.userService.UpdateUser(id, &req)

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully updated."})
}

func (h *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	h.userService.DeleteUser(id)

	return c.JSON(http.StatusOK, map[string]string{"message": "Successfully deleted."})
}
