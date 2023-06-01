package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func main() {
	m := map[string]interface{}{
		"name": "うしお こう",
		"age":  1000,
		"sex":  "♂",
	}

	fmt.Printf("His name is %s\n", m["name"])
	fmt.Printf("His age is  %d\n", m["age"])
	fmt.Printf("His sex is  %s\n", m["sex"])
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", hello)
	e.GET("/user/:id", GetId)

	e.GET("/usio", GetQueryParam)
	e.GET("/putTest/:id", UpdateUsers)
	e.GET("/getUserTest", GetUsers)
	e.Logger.Fatal(e.Start(":3939"))

}

// パスパラメータを受け取って返すらしい
func GetId(c echo.Context) error {
	// ここでパスパラメータを受け取ってる
	id := "うしおこう" //c.Param("id")
	return c.String(http.StatusOK, id)
}

func GetQueryParam(c echo.Context) error {

	//クエリパラメータの取得
	sort := c.QueryParam("sort")
	limit := c.QueryParam("limit")

	//map型 map[キーの型名]値の型名{キー:値, キー:値, ...}
	//interface{}型はint, float, string などさまざまな型の値を格納できるが、演算はできない
	res := map[string]interface{}{
		"sort":  sort,
		"limit": limit,
	}

	//JSONで返す
	return c.JSON(http.StatusOK, res)
}
// ヘッダーのコンテンツタイプを受け取って返す
func UpdateUsers(c echo.Context) error {
	return c.String(http.StatusOK, c.Request().Header.Get("Content-Type"))
}

// クエリパラメータを受け取って返す
func GetUsers(c echo.Context) error {

	u := request.GetUsers{}
	if err := c.Bind(&u); err != nil {
		err = errors.Wrap(err, "failed to bind request data.")
	}

	return c.JSON(http.StatusOK, err)
}