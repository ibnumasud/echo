package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", getHelloWorld)
	e.GET("/user/:id", getUser)
	e.Logger.Fatal(e.Start(":1323"))
}

func getHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Ibnu dari Echo")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
