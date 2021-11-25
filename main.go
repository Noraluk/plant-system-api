package main

import (
	"os"
	"plant-system-api/api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	route := routes.NewRoute(e)
	route.Init()

	e.Logger.Fatal(e.Start(os.Getenv("PORT")))
}
