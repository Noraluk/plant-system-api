package pumpctrlitf

import "github.com/labstack/echo/v4"

type PumpController interface {
	ActivePump(c echo.Context) error
}
