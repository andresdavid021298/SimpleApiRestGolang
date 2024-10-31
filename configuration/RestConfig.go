package configuration

import (
	"ApiRest/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RestConfig() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	port := util.LoadProperty("SERVER_PORT")
	e.Logger.Fatal(e.Start((":" + port)))
}
