package pumpserviceitf

import pumpmodel "plant-system-api/api/models/pump"

type PumpService interface {
	ActivePump(pump *pumpmodel.Pump) error
	AskPump(pump *pumpmodel.Pump) error

	GetPump(id int) (*pumpmodel.Pump, error)
}
