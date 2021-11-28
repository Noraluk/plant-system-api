package config

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/db"
	"google.golang.org/api/option"
)

func NewFirebaseClient(conf *firebase.Config, opt option.ClientOption) (*db.Client, error) {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, conf, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app : %w", err)
	}

	client, err := app.Database(ctx)
	if err != nil {
		return nil, fmt.Errorf("error initializing database client : %w", err)
	}
	return client, nil
}

type Client interface {
	NewRef(path string) Ref
}

type client struct {
	*db.Client
}

func NewClient(firebaseClient *db.Client) Client {
	return &client{firebaseClient}
}

func (conf *client) NewRef(path string) Ref {
	return &ref{conf.Client.NewRef(path)}
}

type Ref interface {
	Child(path string) Ref
	Set(ctx context.Context, v interface{}) error
}

type ref struct {
	*db.Ref
}

func (conf *ref) Child(path string) Ref {
	return &ref{conf.Ref.Child(path)}
}

func (conf *ref) Set(ctx context.Context, v interface{}) error {
	return conf.Ref.Set(ctx, v)
}
