package main

import (
"net/http"

"github.com/labstack/echo"
"github.com/labstack/echo/middleware"
)

func getUser(c echo.Context) error {
   code := c.Param("code")
   return c.String(http.StatusServiceUnavailable, code)
}

func main() {
  e := echo.New()
  e.Use(middleware.Logger())

  e.GET("/status/:code", getUser)

  e.Logger.Fatal(e.Start(":1323"))
}
