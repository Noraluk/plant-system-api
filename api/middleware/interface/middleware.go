package middlewareitf

import "github.com/labstack/echo/v4"

type Middleware interface {
	AskPump(next echo.HandlerFunc) echo.HandlerFunc
}
