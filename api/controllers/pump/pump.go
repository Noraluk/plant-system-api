package pumpctrl

import (
	"log"
	"net/http"
	pumpctrlitf "plant-system-api/api/controllers/pump/interface"
	pumpmodel "plant-system-api/api/models/pump"
	pumpservice "plant-system-api/api/services/pump"
	pumpserviceitf "plant-system-api/api/services/pump/interface"
	"plant-system-api/config"

	"github.com/labstack/echo/v4"
)

type pumpController struct {
	pumpService pumpserviceitf.PumpService
}

func NewPumpController(firebaseClient config.Client) pumpctrlitf.PumpController {
	return &pumpController{
		pumpService: pumpservice.NewPumpService(firebaseClient),
	}
}

func (ct *pumpController) ActivePump(c echo.Context) error {
	req := new(pumpmodel.PumpActiveReq)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pump := &pumpmodel.Pump{ID: req.ID, IsActive: req.IsActive}
	err := ct.pumpService.ActivePump(pump)
	if err != nil {
		log.Println("active pump error : ", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pumpmodel.PumpActiveResponse{ID: pump.ID, IsActive: pump.IsActive})
}

func (ct *pumpController) IsPumpWorking(c echo.Context) error {
	req := new(pumpmodel.PumpIsWorkingReq)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pump := &pumpmodel.Pump{ID: req.ID, IsWorking: req.IsWorking}
	err := ct.pumpService.IsPumpWorking(pump)
	if err != nil {
		log.Println("checking pump working error : ", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pumpmodel.PumpIsWorkingResponse{ID: pump.ID, IsWorking: pump.IsWorking})
}
