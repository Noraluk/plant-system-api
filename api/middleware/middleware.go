package middleware

import (
	"errors"
	"log"
	middlewareitf "plant-system-api/api/middleware/interface"
	pumpmodel "plant-system-api/api/models/pump"
	pumpservice "plant-system-api/api/services/pump"
	pumpserviceitf "plant-system-api/api/services/pump/interface"
	"plant-system-api/config"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type middleware struct {
	pumpService pumpserviceitf.PumpService
}

func NewMiddleware(firebaseClient config.Client) middlewareitf.Middleware {
	return &middleware{
		pumpService: pumpservice.NewPumpService(firebaseClient),
	}
}

func (m *middleware) AskPump(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.Error(err)
		}

		pump := &pumpmodel.Pump{ID: id, IsAsk: true, IsWorking: false}
		err = m.pumpService.AskPump(pump)
		if err != nil {
			log.Println("asking pump working error : ", err)
			c.Error(errors.New("internal server error"))
		}

		time.Sleep(2 * time.Second)
		pump, err = m.pumpService.GetPump(id)
		if err != nil {
			log.Println("getting pump working error : ", err)
			c.Error(errors.New("internal server error"))
		}

		if !pump.IsWorking {
			c.Error(errors.New("pump is not working"))
		}

		return nil
	}
}
