package configuration

import (
	"ApiRest/controller"
	"ApiRest/util"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

func RestConfig() {
	e := echo.New()
	e.Validator = &util.RequestRestValidator{Validator: validator.New()}
	port := util.LoadProperty("SERVER_PORT")
	setControllers(e)
	e.Logger.Fatal(e.Start((":" + port)))
}

func setControllers(echo *echo.Echo) {
	motorcycleController := controller.NewMotorcycleController(echo)
	motorcycleController.AddEndpoints()

	carController := controller.NewCarController(echo)
	carController.AddEndpoints()
}
