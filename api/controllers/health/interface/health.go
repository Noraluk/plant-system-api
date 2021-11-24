package healthctlritf

import "github.com/labstack/echo/v4"

type HealthController interface {
	GetHealth(c echo.Context) error
}
