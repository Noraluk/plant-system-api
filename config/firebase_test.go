package config

import (
	"context"
	"testing"

	firebase "firebase.google.com/go/v4"
	"github.com/stretchr/testify/assert"
	"google.golang.org/api/option"
)

func TestNewFirebaseClient(t *testing.T) {
	const (
		mockDB = "https://mock-db.firebaseio.com"
	)
	type arg struct {
		conf *firebase.Config
		opt  option.ClientOption
	}
	tests := []struct {
		name    string
		arg     arg
		wantErr bool
	}{
		{
			name: "success",
			arg: arg{
				conf: &firebase.Config{
					DatabaseURL: mockDB,
				},
				opt: option.WithoutAuthentication(),
			},
			wantErr: false,
		},
		{
			name: "error without database url",
			arg: arg{
				conf: &firebase.Config{},
				opt:  option.WithoutAuthentication(),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFirebaseClient(tt.arg.conf, tt.arg.opt)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFirebaseClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				assert.NotNil(t, got)
			}

		})
	}
}

func Test_client_NewRef(t *testing.T) {
	const (
		mockDB = "https://mock-db.firebaseio.com"
	)
	var (
		conf = &firebase.Config{
			DatabaseURL: mockDB,
		}
		opt = option.WithoutAuthentication()
	)

	firebaseClient, err := NewFirebaseClient(conf, opt)
	assert.NoError(t, err)

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				path: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewClient(firebaseClient)
			got := client.NewRef(tt.args.path)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_ref_Child(t *testing.T) {
	const (
		mockDB = "https://mock-db.firebaseio.com"
	)
	var (
		conf = &firebase.Config{
			DatabaseURL: mockDB,
		}
		opt = option.WithoutAuthentication()
	)

	firebaseClient, err := NewFirebaseClient(conf, opt)
	assert.NoError(t, err)
	client := NewClient(firebaseClient)

	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				path: "",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ref := client.NewRef(tt.args.path)
			got := ref.Child(tt.args.path)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_ref_Set(t *testing.T) {
	const (
		mockDB = "https://mock-db.firebaseio.com"
	)
	var (
		conf = &firebase.Config{
			DatabaseURL: mockDB,
		}
		opt = option.WithoutAuthentication()
	)

	firebaseClient, err := NewFirebaseClient(conf, opt)
	assert.NoError(t, err)
	client := NewClient(firebaseClient)

	type args struct {
		ctx context.Context
		v   interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "cannot set by mockdb (fake db)",
			args: args{
				ctx: context.Background(),
				v:   "",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ref := client.NewRef("")
			err := ref.Set(tt.args.ctx, tt.args.v)
			if tt.wantErr {
				assert.Error(t, err)
			}
		})
	}
}
