package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/user/:id", GetId)
	e.Logger.Fatal(e.Start(":3939"))
}

// パスパラメータを受け取って返すらしい
func GetId(c echo.Context) error {
	// ここでパスパラメータを受け取ってる
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
