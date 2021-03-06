package pumpservice

import (
	pumpmodel "plant-system-api/api/models/pump"
	pumprepo "plant-system-api/api/repositories/pump"
	pumprepoitf "plant-system-api/api/repositories/pump/interface"
	pumpserviceitf "plant-system-api/api/services/pump/interface"
	"plant-system-api/config"
)

type pumpService struct {
	pumpRepository pumprepoitf.PumpRepository
}

func NewPumpService(firebaseClient config.Client) pumpserviceitf.PumpService {
	return &pumpService{
		pumpRepository: pumprepo.NewPumpRepository(firebaseClient),
	}
}

func (s *pumpService) ActivePump(pump *pumpmodel.Pump) error {
	return s.pumpRepository.ActivePump(pump)
}

func (s *pumpService) AskPump(pump *pumpmodel.Pump) error {
	err := s.pumpRepository.IsPumpWorking(pump)
	if err != nil {
		return err
	}

	return s.pumpRepository.AskPump(pump)
}

func (s *pumpService) GetPump(id int) (*pumpmodel.Pump, error) {
	return s.pumpRepository.GetPump(&pumpmodel.Pump{ID: id})
}
