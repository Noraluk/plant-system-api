package routes

import (
	"firebase.google.com/go/v4/db"
	"github.com/labstack/echo/v4"

	healthrt "plant-system-api/api/routes/health"
	healthrtitf "plant-system-api/api/routes/health/interface"
	pumprt "plant-system-api/api/routes/pump"
	pumprtitf "plant-system-api/api/routes/pump/interface"
)

type Route interface {
	Init()
}

type route struct {
	healthRoute healthrtitf.HealthRoute
	pumpRoute   pumprtitf.PumpRoute
}

func NewRoute(e *echo.Echo, firebaseClient *db.Client) Route {
	return &route{
		healthRoute: healthrt.NewHealthRoute(e),
		pumpRoute:   pumprt.NewPumpRoute(e, firebaseClient),
	}
}

func (r *route) Init() {
	r.healthRoute.SetRoutes()
	r.pumpRoute.SetRoutes()
}
