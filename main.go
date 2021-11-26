package main

import (
	"fmt"
	"os"
	"plant-system-api/api/routes"
	"plant-system-api/config"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	firebaseClient, err := config.NewFirebaseClient()
	if err != nil {
		e.Logger.Fatal(err)
	}

	route := routes.NewRoute(e, firebaseClient)
	route.Init()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
