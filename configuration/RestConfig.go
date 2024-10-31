package configuration

import (
	"ApiRest/controller"
	"ApiRest/util"

	"github.com/labstack/echo/v4"
)

func RestConfig() {
	e := echo.New()
	port := util.LoadProperty("SERVER_PORT")
	setControllers(e)
	e.Logger.Fatal(e.Start((":" + port)))
}

func setControllers(echo *echo.Echo) {
	// Add controller motorcycle
	motorcycleController := controller.NewMotorcycleController(echo)
	motorcycleController.AddEndpoints()

	// // Add controller car
	carController := controller.NewCarController(echo)
	carController.AddEndpoints()
}
