package pumpservice

import (
	"firebase.google.com/go/v4/db"

	pumpmodel "plant-system-api/api/models/pump"
	pumprepo "plant-system-api/api/repositories/pump"
	pumprepoitf "plant-system-api/api/repositories/pump/interface"
	pumpserviceitf "plant-system-api/api/services/pump/interface"
)

type pumpService struct {
	pumpRepository pumprepoitf.PumpRepository
}

func NewPumpService(firebaseClient *db.Client) pumpserviceitf.PumpService {
	return &pumpService{
		pumpRepository: pumprepo.NewPumpRepository(firebaseClient),
	}
}

func (s *pumpService) ActivePump(pump *pumpmodel.Pump) error {
	return s.pumpRepository.ActivePump(pump)
}
