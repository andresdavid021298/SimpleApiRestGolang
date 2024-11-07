package controller

import (
	"ApiRest/request"
	"ApiRest/service"
	"ApiRest/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CarController struct {
	echo       *echo.Echo
	carService *service.CarService
}

func NewCarController(echo *echo.Echo) *CarController {
	return &CarController{echo: echo, carService: service.NewCarService()}
}

func (controller *CarController) AddEndpoints() {
	carGroup := controller.echo.Group("/car")
	carGroup.GET("", controller.greeting)
	carGroup.GET("/all", controller.getAllCars)
	carGroup.GET("/byBrand/:brand", controller.getCarsByBrand)
	carGroup.GET("/used", controller.getUsedCars)
	carGroup.DELETE("/:id", controller.deleteCar)
	carGroup.POST("", controller.insert)
}

func (controller *CarController) greeting(echoContext echo.Context) error {
	greeting := "====================\n" + "WELCOME CAR MODULE" + "\n===================="
	return echoContext.String(http.StatusOK, greeting)
}

func (controller *CarController) getAllCars(echoContext echo.Context) error {
	allCars := controller.carService.GetAll()
	return echoContext.JSON(http.StatusOK, allCars)
}

func (controller *CarController) getCarsByBrand(echoContext echo.Context) error {
	brand := echoContext.Param("brand")
	carsByBrand := controller.carService.GetByBrand(brand)
	return echoContext.JSON(http.StatusOK, carsByBrand)
}

func (controller *CarController) getUsedCars(echoContext echo.Context) error {
	carsUsed := controller.carService.GetUsed()
	return echoContext.JSON(http.StatusOK, carsUsed)
}

func (controller *CarController) insert(echoContext echo.Context) error {
	var newCarRequest request.NewCarRequest
	errorBind := echoContext.Bind(&newCarRequest)
	errorValidate := echoContext.Validate(newCarRequest)
	if errorMessage := util.ValidateRequestErrors(errorBind, errorValidate); errorMessage != "" {
		return echoContext.String(http.StatusBadRequest, errorMessage)
	}
	response := controller.carService.InsertNewCar(newCarRequest)
	return echoContext.String(http.StatusOK, response)
}

func (controller *CarController) deleteCar(echoContext echo.Context) error {
	id := echoContext.Param("id")
	response := controller.carService.DeleteCar(id)
	return echoContext.String(http.StatusOK, response)
}
