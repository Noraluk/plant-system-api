package pumprt

import (
	pumpctrl "plant-system-api/api/controllers/pump"
	pumpctrlitf "plant-system-api/api/controllers/pump/interface"
	pumprtitf "plant-system-api/api/routes/pump/interface"
	"plant-system-api/config"

	"github.com/labstack/echo/v4"
)

type pumpRoute struct {
	e              *echo.Echo
	pumpController pumpctrlitf.PumpController
}

func NewPumpRoute(e *echo.Echo, firebaseClient config.Client) pumprtitf.PumpRoute {
	return &pumpRoute{
		e:              e,
		pumpController: pumpctrl.NewPumpController(firebaseClient),
	}
}

func (r *pumpRoute) SetRoutes() {
	pumpGroup := r.e.Group("/pumps")
	pumpGroup.PATCH("", r.pumpController.ActivePump)
	pumpGroup.PATCH("/check", r.pumpController.IsPumpWorking)
}
