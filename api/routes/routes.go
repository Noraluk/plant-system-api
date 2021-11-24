package routes

import (
	"github.com/labstack/echo/v4"

	healthrt "plant-system-api/api/routes/health"
	healthrtitf "plant-system-api/api/routes/health/interface"
)

type Route interface {
	Init()
}

type route struct {
	healthRoute healthrtitf.HealthRoute
}

func NewRoute(e *echo.Echo) Route {
	return &route{
		healthRoute: healthrt.NewHealthRoute(e),
	}
}

func (r *route) Init() {
	r.healthRoute.SetRoutes()
}
