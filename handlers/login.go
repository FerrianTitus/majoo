package handlers

import (
	"github.com/majoo/middlewares"
	"github.com/majoo/models"
	"net/http"
	"github.com/majoo/repository"

	"github.com/labstack/echo/v4"
)

//Function untuk melakukan login
func Login(c echo.Context) error {
	var login models.LoginRequest
	if err := c.Bind(&login); err != nil {
		return c.String(http.StatusBadRequest, "Bad Request")
	}

	checkUser, err := repository.GetUserByUsername(login.Username, login.Password)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	if checkUser.ID == 0 {
		return c.String(http.StatusNotFound, "Username or Password Not Found")
	}

	token, err := middlewares.CreateToken(checkUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "Error Generate Token")
	}

	resp := models.LoginResponse{
		Token: token,
	}

	return c.JSON(http.StatusCreated, resp)
}
