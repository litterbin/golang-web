package main

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"net/http"
)

// Handler
func hello(c *echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!\n")
}

func main() {
	e := echo.New()

	//Middleware
	e.Use(mw.Logger())
	e.Use(mw.Recover())

	//Routes
	e.Get("/", hello)

	e.Run(":1234")

}
