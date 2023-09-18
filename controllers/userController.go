package controllers

import (
	"net/http"
	"strconv"

	"github.com/abrarnaim015/KP-golang-tgs7/lib/database"
	"github.com/labstack/echo"
)

func GetUsersController(c echo.Context) error {
	users, err := database.GetUsers()

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "Success get all users",
		"Data": users,
	})
}

func GetUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	user, err := database.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "Success get user",
		"Data": user,
	})
}

func CreateUserController(c echo.Context) error {
	user, err := database.CreateUser(c)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code": 201,
		"msg":  "Success create new user",
		"Data": user,
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	if database.DeleteUser(id) == false {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete user")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "Success delete user",
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid ID")
	}

	_, err = database.GetUser(id)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User not found")
	}

	user, err := database.Updateuser(c, id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code": 200,
		"msg":  "Success update user",
		"Data": user,
	})
}