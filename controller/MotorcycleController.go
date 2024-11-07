package controller

import (
	"ApiRest/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type MotorcycleController struct {
	echo              *echo.Echo
	motorcycleService *service.MotorcycleService
}

func NewMotorcycleController(echo *echo.Echo) *MotorcycleController {
	return &MotorcycleController{echo: echo, motorcycleService: service.NewMotorcycleService()}
}

func (controller *MotorcycleController) AddEndpoints() {
	motorcycleGroup := controller.echo.Group("/motorcycle")
	motorcycleGroup.GET("", controller.greeting)
	motorcycleGroup.GET("/all", controller.getAllMotorcycles)
	motorcycleGroup.GET("/:id", controller.getMotorcycleById)
	motorcycleGroup.GET("/byBrand/:brand", controller.getMotorcyclesByBrand)
	motorcycleGroup.PATCH("/changeAmount", controller.changeAmount)
	motorcycleGroup.DELETE("/:id", controller.deleteMotorcycle)
}

func (controller *MotorcycleController) greeting(echoContext echo.Context) error {
	greeting := "=========================\n" + "WELCOME MOTORCYCLE MODULE" + "\n========================="
	return echoContext.String(http.StatusOK, greeting)
}

func (controller *MotorcycleController) getAllMotorcycles(echoContext echo.Context) error {
	allMotorcycles := controller.motorcycleService.GetAll()
	return echoContext.JSON(http.StatusOK, allMotorcycles)
}

func (controller *MotorcycleController) getMotorcycleById(echoContext echo.Context) error {
	id := echoContext.Param("id")
	motorcycle := controller.motorcycleService.GetById(id)
	if motorcycle == nil {
		return echoContext.String(http.StatusNotFound, "Motorcycle with this ID not found")
	}
	return echoContext.JSON(http.StatusOK, motorcycle)
}

func (controller *MotorcycleController) getMotorcyclesByBrand(echoContext echo.Context) error {
	brand := echoContext.Param("brand")
	motorcyclesByBrand := controller.motorcycleService.GetByBrand(brand)
	return echoContext.JSON(http.StatusOK, motorcyclesByBrand)
}

func (controller *MotorcycleController) changeAmount(echoContext echo.Context) error {
	id := echoContext.QueryParam("id")
	amount := echoContext.QueryParam("amount")
	if id == "" || amount == "" {
		return echoContext.String(http.StatusBadRequest, "Query parameters are incorrect")
	}
	response := controller.motorcycleService.ChangeAmount(id, amount)
	if response == nil {
		return echoContext.String(http.StatusNotFound, "Motorcycle with this ID not found")
	}
	return echoContext.JSON(http.StatusOK, response)
}

func (controller *MotorcycleController) deleteMotorcycle(echoContext echo.Context) error {
	id := echoContext.Param("id")
	response := controller.motorcycleService.DeleteMotorcycle(id)
	return echoContext.String(http.StatusOK, response)
}
