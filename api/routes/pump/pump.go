package pumprt

import (
	pumpctrl "plant-system-api/api/controllers/pump"
	pumpctrlitf "plant-system-api/api/controllers/pump/interface"
	pumprtitf "plant-system-api/api/routes/pump/interface"

	"firebase.google.com/go/v4/db"
	"github.com/labstack/echo/v4"
)

type pumpRoute struct {
	e              *echo.Echo
	pumpController pumpctrlitf.PumpController
}

func NewPumpRoute(e *echo.Echo, firebaseClient *db.Client) pumprtitf.PumpRoute {
	return &pumpRoute{
		e:              e,
		pumpController: pumpctrl.NewPumpController(firebaseClient),
	}
}

func (r *pumpRoute) SetRoutes() {
	pumpGroup := r.e.Group("/pumps")
	pumpGroup.GET("/:id", r.pumpController.ActivePump)
}
