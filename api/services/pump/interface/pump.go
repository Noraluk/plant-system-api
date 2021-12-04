package pumpserviceitf

import pumpmodel "plant-system-api/api/models/pump"

type PumpService interface {
	ActivePump(pump *pumpmodel.Pump) error
	IsPumpWorking(pump *pumpmodel.Pump) error
}
