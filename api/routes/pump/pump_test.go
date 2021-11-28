package pumprt

import (
	pumpctrlitf "plant-system-api/api/controllers/pump/interface"
	mock_pumpctrlitf "plant-system-api/api/controllers/pump/mock"
	mockConfig "plant-system-api/config/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewPumpRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		firebaseClientBehavior func(m *mockConfig.MockClient)
	}
	type args struct {
		e              *echo.Echo
		firebaseClient *mockConfig.MockClient
	}
	tests := []struct {
		name    string
		args    args
		fields  fields
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				e:              echo.New(),
				firebaseClient: mockConfig.NewMockClient(ctrl),
			},
			fields: fields{
				firebaseClientBehavior: func(m *mockConfig.MockClient) {
					m.EXPECT().NewRef(gomock.Any())
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.fields.firebaseClientBehavior(tt.args.firebaseClient)

		t.Run(tt.name, func(t *testing.T) {
			got := NewPumpRoute(tt.args.e, tt.args.firebaseClient)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_pumpRoute_SetRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		e              *echo.Echo
		pumpController pumpctrlitf.PumpController
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				e:              echo.New(),
				pumpController: mock_pumpctrlitf.NewMockPumpController(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &pumpRoute{
				e:              tt.fields.e,
				pumpController: tt.fields.pumpController,
			}
			r.SetRoutes()
		})
	}
}
