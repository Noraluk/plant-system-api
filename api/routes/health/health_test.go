package healthrt

import (
	healthctlritf "plant-system-api/api/controllers/health/interface"
	mock_healthctlritf "plant-system-api/api/controllers/health/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestNewHealthRoute(t *testing.T) {
	e := echo.New()
	newHealthRoute := NewHealthRoute(e)
	assert.NotNil(t, newHealthRoute)
}

func Test_healthRoute_SetRoutes(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type fields struct {
		e                *echo.Echo
		healthController healthctlritf.HealthController
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "success",
			fields: fields{
				e:                echo.New(),
				healthController: mock_healthctlritf.NewMockHealthController(ctrl),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &healthRoute{
				e:                tt.fields.e,
				healthController: tt.fields.healthController,
			}
			r.SetRoutes()
		})
	}
}
