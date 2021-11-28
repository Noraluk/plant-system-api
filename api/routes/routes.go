package routes

import (
	"github.com/labstack/echo/v4"

	healthrt "plant-system-api/api/routes/health"
	healthrtitf "plant-system-api/api/routes/health/interface"
	pumprt "plant-system-api/api/routes/pump"
	pumprtitf "plant-system-api/api/routes/pump/interface"
	"plant-system-api/config"
)

type Route interface {
	Init()
}

type route struct {
	healthRoute healthrtitf.HealthRoute
	pumpRoute   pumprtitf.PumpRoute
}

func NewRoute(e *echo.Echo, firebaseClient config.Client) Route {
	return &route{
		healthRoute: healthrt.NewHealthRoute(e),
		pumpRoute:   pumprt.NewPumpRoute(e, firebaseClient),
	}
}

func (r *route) Init() {
	r.healthRoute.SetRoutes()
	r.pumpRoute.SetRoutes()
}
