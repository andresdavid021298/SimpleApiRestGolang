package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type MotorcycleController struct {
	echo *echo.Echo
}

func NewMotorcycleController(echo *echo.Echo) *MotorcycleController {
	return &MotorcycleController{echo: echo}
}

func (controller *MotorcycleController) AddEndpoints() {
	motorcycleGroup := controller.echo.Group("/motorcycle")
	motorcycleGroup.GET("", controller.greeting)
}

func (controller *MotorcycleController) greeting(echoContext echo.Context) error {
	greeting := "=========================\n" + "WELCOME MOTORCYCLE MODULE" + "\n========================="
	return echoContext.String(http.StatusOK, greeting)
}
