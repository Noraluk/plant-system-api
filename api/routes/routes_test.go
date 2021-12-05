package routes

import (
	mockHealthrtitf "plant-system-api/api/routes/health/mock"
	mockPumprtitf "plant-system-api/api/routes/pump/mock"
	mockConfig "plant-system-api/config/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewRoute(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		e                      *echo.Echo
		firebaseClient         *mockConfig.MockClient
		firebaseClientBehavior func(m *mockConfig.MockClient)
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				firebaseClient: mockConfig.NewMockClient(ctrl),
				firebaseClientBehavior: func(m *mockConfig.MockClient) {
					m.EXPECT().NewRef(gomock.Any()).Times(2)
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt.args.firebaseClientBehavior(tt.args.firebaseClient)

		t.Run(tt.name, func(t *testing.T) {
			got := NewRoute(tt.args.e, tt.args.firebaseClient)
			if !tt.wantErr {
				assert.NotNil(t, got)
			}
		})
	}
}

func Test_route_Init(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		healthRoute         *mockHealthrtitf.MockHealthRoute
		healthRouteBehavior func(m *mockHealthrtitf.MockHealthRoute)
		pumpRoute           *mockPumprtitf.MockPumpRoute
		pumpRouteBehavior   func(m *mockPumprtitf.MockPumpRoute)
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				healthRoute: mockHealthrtitf.NewMockHealthRoute(ctrl),
				healthRouteBehavior: func(m *mockHealthrtitf.MockHealthRoute) {
					m.EXPECT().SetRoutes()
				},
				pumpRoute: mockPumprtitf.NewMockPumpRoute(ctrl),
				pumpRouteBehavior: func(m *mockPumprtitf.MockPumpRoute) {
					m.EXPECT().SetRoutes()
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.healthRouteBehavior(tt.fields.healthRoute)
			tt.fields.pumpRouteBehavior(tt.fields.pumpRoute)

			r := &route{
				healthRoute: tt.fields.healthRoute,
				pumpRoute:   tt.fields.pumpRoute,
			}
			r.Init()
		})
	}
}
