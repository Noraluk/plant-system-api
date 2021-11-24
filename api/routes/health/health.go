package healthrt

import (
	healthctlr "plant-system-api/api/controllers/health"
	healthctlritf "plant-system-api/api/controllers/health/interface"
	healthrtitf "plant-system-api/api/routes/health/interface"

	"github.com/labstack/echo/v4"
)

type healthRoute struct {
	e                *echo.Echo
	healthController healthctlritf.HealthController
}

func NewHealthRoute(e *echo.Echo) healthrtitf.HealthRoute {
	return &healthRoute{
		e:                e,
		healthController: healthctlr.NewHealthController(),
	}
}

func (r *healthRoute) SetRoutes() {
	r.e.GET("/health", r.healthController.GetHealth)
}
