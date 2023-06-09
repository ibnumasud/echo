package main

import (
	"database/sql"
	"fmt"
	"myapp/controllers"
	"myapp/sqldb"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apmecho "go.elastic.co/apm/module/apmechov4/v2"
	"go.elastic.co/apm/v2"
)

func main() {
	e := echo.New()
	//Middleware
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {}))
	//CORS curl -H "Origin: https://masud.id"
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://masud.id", "https://masud.id"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	//CRSF yang ini susah ditest
	e.Use(middleware.CSRF())

	//APM Elastic
	// e.Use(apmecho.Middleware())
	db := sqldb.ConnectDB()

	h := controllers.NewBaseHandler(db)
	if hello, err := h.HelloWorld(); err != nil {
		fmt.Println("error db")
	}

	e.GET("/", h.HelloWorld())
	e.GET("/user/:id", getUser, apmecho.Middleware())
	e.POST("/users", postUsers)
	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/echo")
	if err != nil {
		return err
	}
	defer db.Close()

	span, _ := apm.StartSpan(c.Request().Context(), "work", "custom")

	//single query
	var nama string
	erro := db.QueryRow("select nama from users where id=?", id).Scan(&nama)
	if erro != nil {
		panic(erro.Error())
	}

	defer span.End()

	return c.String(http.StatusOK, nama)
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
	return c.JSON(http.StatusOK, u)
}
