package main

import (
	"net/http"

	"github.com/labstack/echo"
	"strconv"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"msg": "Hello, World!"})
	})
	e.GET("/users/:id/:age", func(c echo.Context) error {
		age, err := strconv.Atoi(c.Param("age"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]interface{}{"id": c.Param("id"), "age": age})
	})

	type User struct {
		Name string `json:"name"`
	}

	e.POST("/users", func(c echo.Context) error {
		m := new(map[string]interface{})
		if err := c.Bind(m); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, m)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
