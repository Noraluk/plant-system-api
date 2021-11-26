package pumprepoitf

import (
	pumpmodel "plant-system-api/api/models/pump"
)

type PumpRepository interface {
	ActivePump(pump *pumpmodel.Pump) error
}
