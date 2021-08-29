package lessons

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GoModules() {
	// instanciar echo
	e := echo.New()

	// Ruta
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
