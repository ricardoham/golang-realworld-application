package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/ricardoham/pokedex-api/config"
)

func main() {
	apiConfig := config.GetAPIConfig()
	echo := echo.New()

	echo.Use(middleware.Logger())
	echo.Use(middleware.Recover())

	port := fmt.Sprintf(":%d", apiConfig.HostPort)
	echo.Logger.Fatal(echo.Start(port))
}
