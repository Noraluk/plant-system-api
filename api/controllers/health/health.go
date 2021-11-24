package healthctlr

import (
	"net/http"
	healthctlritf "plant-system-api/api/controllers/health/interface"

	"github.com/labstack/echo/v4"
)

type healthController struct{}

func NewHealthController() healthctlritf.HealthController {
	return &healthController{}
}

func (ct *healthController) GetHealth(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
}
