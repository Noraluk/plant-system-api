package pumprt

import (
	pumpctrl "plant-system-api/api/controllers/pump"
	pumpctrlitf "plant-system-api/api/controllers/pump/interface"
	"plant-system-api/api/middleware"
	middlewareitf "plant-system-api/api/middleware/interface"
	pumprtitf "plant-system-api/api/routes/pump/interface"
	"plant-system-api/config"

	"github.com/labstack/echo/v4"
)

type pumpRoute struct {
	e              *echo.Echo
	pumpController pumpctrlitf.PumpController
	middleware     middlewareitf.Middleware
}

func NewPumpRoute(e *echo.Echo, firebaseClient config.Client) pumprtitf.PumpRoute {
	return &pumpRoute{
		e:              e,
		pumpController: pumpctrl.NewPumpController(firebaseClient),
		middleware:     middleware.NewMiddleware(firebaseClient),
	}
}

func (r *pumpRoute) SetRoutes() {
	pumpGroup := r.e.Group("/pumps/:id")
	pumpGroup.Use(r.middleware.AskPump)

	pumpGroup.PATCH("/active", r.pumpController.ActivePump)
	pumpGroup.PATCH("/ask", r.pumpController.AskPump)

	pumpGroup.GET("", r.pumpController.GetPump)
}
