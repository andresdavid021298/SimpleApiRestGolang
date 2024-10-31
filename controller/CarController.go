package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type CarController struct {
	echo *echo.Echo
}

func NewCarController(echo *echo.Echo) *CarController {
	return &CarController{echo: echo}
}

func (controller *CarController) AddEndpoints() {
	carGroup := controller.echo.Group("/car")
	carGroup.GET("", controller.greeting)
}

func (controller *CarController) greeting(echoContext echo.Context) error {
	greeting := "====================\n" + "WELCOME CAR MODULE" + "\n===================="
	return echoContext.String(http.StatusOK, greeting)
}
