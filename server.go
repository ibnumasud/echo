package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", getHelloWorld)
	e.GET("/user/:id", getUser)
	e.POST("/users", postUsers)
	e.Logger.Fatal(e.Start(":1323"))
}

func getHelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Ibnu dari Echo")
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

type User struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Email string `json:"email" xml:"email" form:"email" query:"email"`
}

func postUsers(c echo.Context) error {
	u := new(User)
	//Assignment statement, err nya di assign disini aja
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.XML(http.StatusOK, u)
}
