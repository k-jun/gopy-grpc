package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type simpleResponse struct {
	Message string `json: "message"`
}

func HelloEcho(c echo.Context) error {
	return c.JSON(http.StatusOK, simpleResponse{Message: "this api is working! Ver 1.0.0"})
}
