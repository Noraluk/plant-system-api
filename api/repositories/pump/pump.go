package pumprepo

import (
	"context"
	"fmt"
	pumpmodel "plant-system-api/api/models/pump"
	pumprepoitf "plant-system-api/api/repositories/pump/interface"
	"plant-system-api/config"
)

type pumpRepository struct {
	ctx context.Context
	ref config.Ref
}

func NewPumpRepository(firebaseClient config.Client) pumprepoitf.PumpRepository {
	return &pumpRepository{
		ctx: context.Background(),
		ref: firebaseClient.NewRef("pump"),
	}
}

func (repo *pumpRepository) ActivePump(pump *pumpmodel.Pump) error {
	specificPump := repo.ref.Child(fmt.Sprintf("%d/is_active", pump.ID))
	if err := specificPump.Set(repo.ctx, pump.IsActive); err != nil {
		return fmt.Errorf("error activating pump : %w", err)
	}
	return nil
}

func (repo *pumpRepository) AskPump(pump *pumpmodel.Pump) error {
	specificPump := repo.ref.Child(fmt.Sprintf("%d/is_ask", pump.ID))
	if err := specificPump.Set(repo.ctx, pump.IsAsk); err != nil {
		return fmt.Errorf("error ask pump : %w", err)
	}
	return nil
}

func (repo *pumpRepository) IsPumpWorking(pump *pumpmodel.Pump) error {
	specificPump := repo.ref.Child(fmt.Sprintf("%d/is_working", pump.ID))
	if err := specificPump.Set(repo.ctx, pump.IsWorking); err != nil {
		return fmt.Errorf("error is pump working : %w", err)
	}
	return nil
}

func (repo *pumpRepository) GetPump(id int) (*pumpmodel.Pump, error) {
	var pump *pumpmodel.Pump
	specificPump := repo.ref.Child(fmt.Sprintf("%d", id))
	if err := specificPump.Get(repo.ctx, &pump); err != nil {
		return nil, fmt.Errorf("error getting pump working : %w", err)
	}
	pump.ID = id
	return pump, nil
}
