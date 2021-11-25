package routes

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	mockHealthrtitf "plant-system-api/api/routes/health/mock"
)

func TestNewRoute(t *testing.T) {
	e := echo.New()
	newRoute := NewRoute(e)
	assert.NotNil(t, newRoute)
}

func Test_route_Init(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		healthRoute         *mockHealthrtitf.MockHealthRoute
		healthRouteBehavior func(m *mockHealthrtitf.MockHealthRoute)
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.healthRouteBehavior(tt.fields.healthRoute)

			r := &route{
				healthRoute: tt.fields.healthRoute,
			}
			r.Init()
		})
	}
}
