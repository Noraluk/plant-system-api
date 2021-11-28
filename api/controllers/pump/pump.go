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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	isActive, err := strconv.ParseBool(c.QueryParam("isActive"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	pump := &pumpmodel.Pump{PumpID: id, IsActive: isActive}
	err = ct.pumpService.ActivePump(pump)
	if err != nil {
		log.Println("active pump error : ", err)
		return c.JSON(http.StatusInternalServerError, "internal server error")
	}

	return c.JSON(http.StatusOK, pump)
}
