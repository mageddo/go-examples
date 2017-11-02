package main

import (
	"net/http"

	"github.com/labstack/echo"
	"strconv"
	"github.com/mageddo/go-logging"
	"github.com/labstack/gommon/log"
	"context"

	"io"
	"fmt"
	"crypto/rand"
)

func main() {
	e := echo.New()
	//out, _ := os.Open(os.DevNull)
	//e.Logger.SetOutput(out)
	e.Logger.SetLevel(log.ERROR)
	//e.Use(middleware.Gzip())

	// custom middleware to set request id
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			var err error
			var uuid string
			if uuid, err = newUUID(); err != nil {
				return err
			}
			c.SetRequest(c.Request().WithContext(context.WithValue(c.Request().Context(), "id", uuid)))
			return next(c)
		}
	})

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
		logging.Infof("ctx=%v", c.Request().Context().Value("id"))
		m := new(map[string]interface{})
		if err := c.Bind(m); err != nil {
			return err
		}
		return c.JSON(http.StatusOK, m)
	})
	e.Logger.Fatal(e.Start(":1323"))
}


func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%X-%X-%X-%X-%X", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}