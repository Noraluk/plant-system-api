package main

import (
	"fmt"
	"log"
	"os"
	"plant-system-api/api/routes"
	"plant-system-api/config"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func main() {
	e := echo.New()

	conf := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_DB_URL"),
	}

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_FILE_PATH"))
	firebaseClient, err := config.NewFirebaseClient(conf, opt)
	if err != nil {
		// e.Logger.Fatal(err)
		log.Println(err)
	}

	route := routes.NewRoute(e, config.NewClient(firebaseClient))
	route.Init()

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
