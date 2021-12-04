package pumpctrl

import (
	"log"
	"net/http"
	pumpctrlitf "plant-system-api/api/controllers/pump/interface"
	pumpmodel "plant-system-api/api/models/pump"
	pumpservice "plant-system-api/api/services/pump"
	pumpserviceitf "plant-system-api/api/services/pump/interface"
	"plant-system-api/config"
	"strconv"

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

func (ct *pumpController) AskPump(c echo.Context) error {
	req := new(pumpmodel.PumpAskingReq)
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	pump := &pumpmodel.Pump{ID: req.ID, IsAsk: req.IsAsk}
	err := ct.pumpService.AskPump(pump)
	if err != nil {
		log.Println("asking pump working error : ", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pumpmodel.PumpAskingReq{ID: pump.ID, IsAsk: pump.IsAsk})
}

func (ct *pumpController) GetPump(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pump, err := ct.pumpService.GetPump(id)
	if err != nil {
		log.Println("getting pump working error : ", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pump)
}
