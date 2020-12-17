package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Start")
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	port := fmt.Sprintf(":%d", 8080)
	echo.Logger.Fatal(echo.Start(port))
}
