package config

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func NewFirebaseClient() (*db.Client, error) {
	ctx := context.Background()
	conf := &firebase.Config{
		DatabaseURL: os.Getenv("FIREBASE_DB_URL"),
	}

	opt := option.WithCredentialsFile(os.Getenv("FIREBASE_CREDENTIALS_FILE_PATH"))
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		// log.Fatalln("Error initializing app:", err)
		return nil, fmt.Errorf("error initializing app : %w", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing database client : %w", err)
	}
	return client, nil
}
